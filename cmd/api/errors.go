package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) invalidAuthenticationTokenResponse(c *gin.Context) {
	c.Writer.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message})
}

func (app *application) authenticationRequiredResponse(c *gin.Context) {
	message := "you must be authenticated to access this resource"
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": message})
}

func (app *application) inactiveAccountResponse(c *gin.Context) {
	message := "your user account must be activated to access this resource"

	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": message})
}

func (app *application) rateLimitExceededResponse(c *gin.Context) {
	message := "rate limit exceeded"
	c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": message})
}

func (app *application) notFoundResponse(c *gin.Context) {
	message := "the requested resource could not be found"
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": message})
}

func (app *application) internalServerError(c *gin.Context, message string) {
	if message == "" {
		message = "the server encountered an unexpected error"
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": message})
}

func (app *application) badRequestResponse(c *gin.Context, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message, "data": data})
}
