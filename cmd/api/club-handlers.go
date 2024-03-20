package main

import (
	"ctfrancia/ajedrez-be/internal/data"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createNewClub(c *gin.Context) {
	var cnc data.Club
	if err := c.ShouldBindJSON(&cnc); err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("CREATE NEW CLUB", cnc)
	c.JSON(http.StatusOK, gin.H{
		"message": "create new user",
	})

}
