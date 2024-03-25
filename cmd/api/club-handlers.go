package main

import (
	"ctfrancia/ajedrez-be/internal/data"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createNewClub(c *gin.Context) {
	var cnc data.Club
	if err := c.ShouldBindJSON(&cnc); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), cnc)
		return
	}

	err := app.models.Clubs.Insert(&cnc)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"club_name_unique_check\"" {
			apiResponse(c, http.StatusBadRequest, "error", "Club already exists", cnc)
			return
		}

		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), cnc)
		return
	}

	apiResponse(c, http.StatusCreated, "success", "Club created", cnc)
}

func (app *application) getClubByName(c *gin.Context) {
	var club data.Club
	name := c.Param("name")
	if err := c.ShouldBindUri(&name); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), club)
		return
	}

	err := app.models.Clubs.GetByName(name, &club)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			apiResponse(c, http.StatusNotFound, "error", "Club not found", club)
			return
		}
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), club)
		return
	}

	apiResponse(c, http.StatusOK, "success", "Club found", club)
}
