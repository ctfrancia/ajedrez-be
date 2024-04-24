package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/lib/pq"
	// "database/sql"
	// "io"
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

	fmt.Println("Creating new tournament", input)
	tCode := uuid.New().String()
	t := data.Tournament{
		Name: &input.Name,
		Code: &tCode,
	}

	err := app.models.Tournaments.Insert(&t)
	if err != nil {
		app.internalServerError(c, err.Error())
		return
	}

	c.Writer.Header().Set("Location", "/tournament/"+tCode)
	apiResponse(c, http.StatusCreated, "success", "", t)
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

// can add/remove players in a tournament
func (app *application) updatePlayers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// can add/remove teams in a tournament
func (app *application) updateTeams(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) modifyPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) updateTournament(c *gin.Context) {
	var ut map[string]interface{}
	t := c.MustGet("input").(data.Tournament)

	if t.Code == nil {
		app.badRequestResponse(c, "code is required", nil)
		return
	}

	_, err := uuid.Parse(*t.Code)
	if err != nil {
		data := map[string]interface{}{"code": *t.Code}
		app.badRequestResponse(c, "invalid tournament code", data)
		return
	}

	ut = prepareTournamentUpdate(t)

	err = app.models.Tournaments.Update(ut)
	if err != nil {
		app.internalServerError(c, err.Error())
		return
	}
	c.Writer.Header().Set("Location", "/tournament/"+*t.Code)
	apiResponse(c, http.StatusNoContent, "success", "", ut)
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
