package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) healthcheck(c *gin.Context) {
	systemInfo := gin.H{
		"status": "available",
		"system_info": map[string]interface{}{
			"environment": app.config.env,
			"version":     version,
		},
	}

	c.JSON(http.StatusOK, systemInfo)
}

// TODO: do this below
/*
func (app *application) routes() *gin.Engine {
    r := gin.Default()

    r.Use(app.recoverPanic())
    r.Use(app.rateLimiter(2, 4))
    r.Use(app.setRequestID())
    r.Use(app.logRequest())

    r.POST("/v1/users", app.registerUser)
    r.POST("/v1/tokens/authentication", app.createAuthenticationToken)

    r.GET("/v1/users/me", app.requireAuthenticatedUser(), app.getUser)
    r.PATCH("/v1/users/me", app.requireAuthenticatedUser(), app.updateUser)
    r.DELETE("/v1/users/me", app.requireAuthenticatedUser(), app.deleteUser)

    r.GET("/v1/users/:id", app.requireAuthenticatedUser(), app.getUser)
    r.PATCH("/v1/users/:id", app.requireAuthenticatedUser(), app.updateUser)
    r.DELETE("/v1/users/:id", app.requireAuthenticatedUser(), app.deleteUser)

    return r
}
*/
