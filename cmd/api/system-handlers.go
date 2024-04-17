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
