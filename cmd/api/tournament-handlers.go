package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
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

func (app *application) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.badRequestResponse(c, "invalid tournament id", idStr)
		return
	}

	fmt.Println("Getting tournament by id: ", id)

	t, err := app.models.Tournaments.GetByID(id)
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			app.notFoundResponse(c)

		default:
			app.internalServerError(c, err.Error())
		}

		return
	}

	apiResponse(c, http.StatusOK, "success", "", t)
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

func (app *application) addPlayersToTournament(c *gin.Context) {
	t := c.MustGet("input").(data.Tournament)

	_, err := uuid.Parse(*t.Code)
	if err != nil {
		fmt.Println("error: ", err)
		data := map[string]interface{}{"code": *t.Code}
		app.badRequestResponse(c, "invalid tournament code", data)
		return
	}

	if t.Players == nil {
		app.badRequestResponse(c, "no players to add", t)
		return
	}

	if len(t.Players) == 0 || len(t.Players) > 1 {
		msg := "only one player can be added at a time or there is no player to add"
		app.badRequestResponse(c, msg, t)
		return
	}

	err = app.models.Tournaments.AddPlayersToTournament(t)
	if err != nil {
		switch err {
		case data.ErrNoResultSet:
			c.Writer.Header().Set("Location", "/tournament/"+*t.Code)
			apiResponse(c, http.StatusNoContent, "success", "", t)

		case data.ErrRecordNotFound:
			app.notFoundResponse(c)

		default:
			app.internalServerError(c, err.Error())
		}

		return
	}

	c.Writer.Header().Set("Location", "/tournament/"+*t.Code)
	apiResponse(c, http.StatusNoContent, "success", "", t)
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
	now := time.Now()
	fmt.Println("now: ", now.Format(time.RFC3339))

	_, err := uuid.Parse(*t.Code)
	if err != nil {
		data := map[string]interface{}{"code": *t.Code}
		app.badRequestResponse(c, "invalid tournament code", data)
		return
	}

	ut = prepareTournamentUpdate(t)

	_, err = app.models.Tournaments.Update(ut)
	if err != nil {
		switch err {
		case data.ErrNoResultSet:
			c.Writer.Header().Set("Location", "/tournament/"+*t.Code)
			apiResponse(c, http.StatusNoContent, "success", "", ut)

		case data.ErrRecordNotFound:
			app.notFoundResponse(c)

		default:
			app.internalServerError(c, err.Error())
		}

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
