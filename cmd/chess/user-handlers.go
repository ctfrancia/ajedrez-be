package main

import (
	"ctfrancia/ajedrez-be/internal/models"
	"ctfrancia/ajedrez-be/pkg/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createNewUser(c *gin.Context) {
	var cnu dtos.UserCreateRequest

	if err := c.ShouldBindJSON(&cnu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("cnu", cnu)

	txt := models.CreateUser()
	fmt.Println("txt", txt)

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
	email := c.Param("email")
	resp := gin.H{
		"message": fmt.Sprintf("update user by email: %s", email),
	}

	// response
	c.JSON(http.StatusOK, resp)
}
