package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
func (app *application) invalidAuthenticationTokenResponse(c *gin.Context) {
	c.Writer.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	apiResponse(c, http.StatusUnauthorized, "authentication", "error", message)
	// app.errorResponse(w, r, http.StatusUnauthorized, message)
}
