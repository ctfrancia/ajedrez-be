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
