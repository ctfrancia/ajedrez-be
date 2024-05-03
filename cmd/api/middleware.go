package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/models"
	"errors"
	"expvar"
	"net"
	"net/http"
	"strconv"
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
			c.Set("user", models.AnonymousUser)
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

		user, err := app.repository.Users.GetForToken(data.ScopeAuthentication, token)
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
		user := c.MustGet("user").(*models.User)

		if user.IsAnonymous() {
			app.authenticationRequiredResponse(c)
			return
		}

		c.Next()
	}
}

func (app *application) requireActivatedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.User)

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

type metricsResponseWriter struct {
	wrapped       http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func newMetricsResponseWriter(w http.ResponseWriter) *metricsResponseWriter {
	return &metricsResponseWriter{
		wrapped:    w,
		statusCode: http.StatusOK,
	}
}

func (mw *metricsResponseWriter) Header() http.Header {
	return mw.wrapped.Header()
}

func (mw *metricsResponseWriter) WriteHeader(statusCode int) {
	mw.wrapped.WriteHeader(statusCode)

	if !mw.headerWritten {
		mw.statusCode = statusCode
		mw.headerWritten = true
	}
}

func (mw *metricsResponseWriter) Write(b []byte) (int, error) {
	mw.headerWritten = true
	return mw.wrapped.Write(b)
}

func (mw *metricsResponseWriter) Unwrap() http.ResponseWriter {
	return mw.wrapped
}

func (app *application) metrics() gin.HandlerFunc {
	var (
		totalRequestsReceived           = expvar.NewInt("total_requests_received")
		totalResponsesSent              = expvar.NewInt("total_responses_sent")
		totalProcessingTimeMicroseconds = expvar.NewInt("total_processing_time_μs")
		totalResponsesSentByStatus      = expvar.NewMap("total_responses_sent_by_status")
	)
	return func(c *gin.Context) {

		start := time.Now()
		totalRequestsReceived.Add(1)
		mw := newMetricsResponseWriter(c.Writer)
		c.Next()
		totalResponsesSent.Add(1)

		// TODO: THIS IS NOT WORKING!!!! ALWAYS EQUALS 200!!!
		totalResponsesSentByStatus.Add(strconv.Itoa(mw.statusCode), 1)

		duration := time.Since(start).Microseconds()
		totalProcessingTimeMicroseconds.Add(duration)
	}
}

func (app *application) tournamentValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Tournament
		if err := c.ShouldBindJSON(&input); err != nil {
			msg := "invalid field(s) and/or missing required field(s) in the request body"
			app.badRequestResponse(c, msg, input)
			return
		}

		if input.Code == nil {
			app.badRequestResponse(c, "code is required", input)
			return
		}

		c.Set("tournament", input)
		c.Next()
	}
}

func (app *application) matchValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		   var input data.Match
		   if err := c.ShouldBindJSON(&input); err != nil {
		       msg := "invalid field(s) and/or missing required field(s) in the request body"
		       app.badRequestResponse(c, msg, input)
		       return
		   }

		   if input.TournamentID == nil {
		       app.badRequestResponse(c, "tournament_id is required", input)
		       return
		   }

		   if input.Team1ID == nil {
		       app.badRequestResponse(c, "team1_id is required", input)
		       return
		   }

		   if input.Team2ID == nil {
		       app.badRequestResponse(c, "team2_id is required", input)
		       return
		   }

		   c.Set("input", input)
		*/
		c.Next()
	}
}
