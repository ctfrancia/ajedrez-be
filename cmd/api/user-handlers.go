package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createNewUser(c *gin.Context) {
	var cnu data.User
	if err := c.ShouldBindJSON(&cnu); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), cnu)
		return
	}

	// normalize user data
	normalizeUser(&cnu)

	err := app.models.Users.Insert(&cnu)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_club_user_code_unique\"" {
			apiResponse(c, http.StatusBadRequest, "error", "user already exists", cnu)
			return
		}
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), cnu)
		return
	}

	apiResponse(c, http.StatusCreated, "success", "user created", cnu)
}

func getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	resp := gin.H{
		"message": fmt.Sprintf("get user by email: %s", email),
	}

	// response
	c.JSON(http.StatusOK, resp)
}

func (app *application) getAllUsers(c *gin.Context) {
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
