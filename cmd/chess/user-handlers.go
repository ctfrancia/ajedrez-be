package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createNewUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create new user",
	})
}

func getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	resp := gin.H{
		"message": fmt.Sprintf("get user by email: %s", email),
	}

	// response
	c.JSON(http.StatusOK, resp)
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get users",
	})
}

func updateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update user",
	})
}
