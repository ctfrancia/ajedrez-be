package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func (app *application) createNewTournament(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		app.badRequestResponse(c, err.Error(), input)
		return
	}

	tCode := uuid.New().String()
	t := data.Tournament{
		Name: input.Name,
		Code: tCode,
	}

	err := app.models.Tournaments.Insert(&t)
	if err != nil {
		app.internalServerError(c, err.Error())
		return
	}

	c.Writer.Header().Set("Location", "/tournaments/"+tCode)
	c.JSON(http.StatusCreated, gin.H{
		"message": "no errors",
		"data":    t,
	})
}

func (app *application) getTournament(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// can search by date, name, location, type, rating, etc
func (app *application) getTournaments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) updateTournament(c *gin.Context) {
	var input map[string]interface{}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	json.Unmarshal(jsonData, &input)
	if _, ok := input["code"]; !ok {
		apiResponse(c, http.StatusBadRequest, "error", "code is required", input)
		return
	}

	_, err = uuid.Parse(input["code"].(string))
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", "invalid tournament code", input)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) deleteTournament(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) getTournamentPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) uploadTournamentPoster(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) publishTournament(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
