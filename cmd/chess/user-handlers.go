package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func createNewUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create new user",
	})
}

func getUserByEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get user by email",
	})
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
