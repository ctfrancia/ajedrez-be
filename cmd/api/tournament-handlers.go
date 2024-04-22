package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "io"
	"encoding/json"
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
		Name: input.Name,
		Code: tCode,
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

func (app *application) updateTournament(c *gin.Context) {
	var input map[string]interface{}
	var t data.Tournament
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&t); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	/*
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
			return
		}
	*/

	fmt.Printf("v1 %#v", t)
	fmt.Printf("Updating tournament %#v", t)

	/*

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

		c.Writer.Header().Set("Location", "/tournament/"+input["code"].(string))
	*/
	apiResponse(c, http.StatusNoContent, "success", "", input)
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
