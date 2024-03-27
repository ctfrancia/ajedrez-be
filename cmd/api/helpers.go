package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"github.com/gin-gonic/gin"
	"strings"
)

func apiResponse(c *gin.Context, httpStatus int, status, message string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func normalizeUser(u *data.User) {
	u.Email = strings.ToLower(u.Email)
	u.FirstName = strings.Trim(u.FirstName, " ")
	u.LastName = strings.Trim(u.LastName, " ")
	u.Username = strings.Trim(u.Username, " ")
	u.Country = strings.Trim(u.Country, " ")
}

func prepareUserUpdate(oldData *data.User, newData *data.User) *data.User {

	return &data.User{}
}
