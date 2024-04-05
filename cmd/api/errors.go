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
