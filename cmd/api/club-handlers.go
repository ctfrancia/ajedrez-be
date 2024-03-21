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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Clubs.Insert(&cnc)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"club_name_unique_check\"" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Club already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := gin.H{
		"message": "success",
		"data":    cnc,
	}
	c.JSON(http.StatusCreated, resp)
}

func (app *application) getClubByName(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("name", name)
	club, err := app.models.Clubs.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Club not found"})
		return
	}

	resp := gin.H{
		"message": "success",
		"data":    club,
	}
	c.JSON(http.StatusOK, resp)
}
