package main

import (
	"github.com/gin-gonic/gin"
)

func apiResponse(c *gin.Context, httpStatus int, status, message string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
