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
		fmt.Println("cnu", cnu)
		fmt.Println(" \n error", err.Error())
		resp := gin.H{
			"status":  "error",
			"message": err.Error(),
			"user":    cnu,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := app.models.Users.Insert(&cnu)
	if err != nil {
        if err.Error() == "pq: duplicate key value violates unique constraint \"users_club_user_code_unique\"" {
            resp := gin.H{
                "status": "error",
                "message": "user already exists",
                "data": cnu,
            }
            c.JSON(http.StatusBadRequest, resp)
            return
        }
        resp := gin.H{
            "status": "error",
            "message": err.Error(),
            "data": cnu,
        }
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
        "status": "success",
		"message": nil,
		"data":    cnu,
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
