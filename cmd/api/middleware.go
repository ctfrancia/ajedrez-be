package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func (app *application) authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Vary", "Authorization")

		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.Set("user", data.AnonymousUser)
			c.Next()
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(c)
			return
		}

		token := headerParts[1]
		err := data.ValidateTokenPlaintext(token)
		if err != nil {
			app.invalidAuthenticationTokenResponse(c)
			return
		}

		user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				apiResponse(c, http.StatusUnauthorized, "Invalid authentication token", "error", nil)
			default:
				apiResponse(c, http.StatusInternalServerError, "Internal server error", "error", nil)
			}
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func (app *application) requireAuthenticatedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*data.User)

		if user.IsAnonymous() {
			app.authenticationRequiredResponse(c)
			return
		}

		c.Next()
	}
}

func (app *application) requireActivatedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*data.User)
		fmt.Println(user)

		if user.IsAnonymous() {
			app.authenticationRequiredResponse(c)
			return
		}

		if !user.Activated {
			app.inactiveAccountResponse(c)
			return
		}

		c.Next()
	}
}

// rateLimit is a middleware function that rate limits the requests to the API based
// on the client's IP address.
func (app *application) rateLimit() gin.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)
	go func() {
		for {
			time.Sleep(time.Minute)
			// Lock the mutex to prevent any rate limiter checks from happening while
			// the cleanup is taking place.
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		if app.config.limiter.enabled {
			ip, _, err := net.SplitHostPort(c.ClientIP())
			if err != nil {
				app.internalServerError(c, err.Error())
				return
			}

			mu.Lock()
			if _, found := clients[ip]; !found {
				clients[ip] = &client{
					limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst),
				}
			}

			// Update the last seen time for the client.
			clients[ip].lastSeen = time.Now()

			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceededResponse(c)
				return
			}

			mu.Unlock()
		}

		c.Next()
	}
}

func (app *application) enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Vary", "Origin")
		c.Writer.Header().Add("Vary", "Access-Control-Request-Method")
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			for i := range app.config.cors.trustedOrigins {
				if origin == app.config.cors.trustedOrigins[i] {
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

					if c.Request.Method == http.MethodOptions && c.GetHeader("Access-Control-Request-Method") != "" {
						c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
						c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
						c.Writer.Header().Set("Access-Control-Max-Age", "86400")
						c.AbortWithStatus(http.StatusOK)
						return
					}

					break
				}
			}
		}

		c.Next()
	}
}
