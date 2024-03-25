package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"strings"
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

func (app *application) getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	email = strings.ToLower(email)
	var user data.User
	err := app.models.Users.GetByEmail(email, &user)
	if err != nil {
		apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user found", user)
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
