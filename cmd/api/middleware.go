package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"strings"
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

func (app *application) rateLimit(reqPerSec rate.Limit, burstReq int) gin.HandlerFunc {
	limiter := rate.NewLimiter(reqPerSec, burstReq)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			apiResponse(c, http.StatusTooManyRequests, "Too many requests", "error", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
