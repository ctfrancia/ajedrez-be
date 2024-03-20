package main

import (
	"ctfrancia/ajedrez-be/internal/data"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createNewClub(c *gin.Context) {
	var cnc data.Club
	if err := c.ShouldBindJSON(&cnc); err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Clubs.Insert(&cnc)
	if err != nil {
		fmt.Println("error", err)
		if err.Error() == "pq: duplicate key value violates unique constraint \"clubs_code_key\"" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Club already exists"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Club already exists"})
		return
	}
	resp := gin.H{
		"message": "success",
		"data":    cnc,
	}
	c.JSON(http.StatusOK, resp)
}
