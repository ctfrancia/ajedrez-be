package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"errors"
	"github.com/gin-gonic/gin"
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
