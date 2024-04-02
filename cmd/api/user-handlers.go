package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"fmt"
	"net/http"

	"io"
	"strings"
	"time"

	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createNewUser(c *gin.Context) {
	// TODO: move this to public dto
	var input struct {
		FirstName        string    `json:"first_name"`
		LastName         string    `json:"last_name"`
		DateOfBirth      time.Time `json:"date_of_birth"`
		Sex              string    `json:"sex"`
		ClubID           int       `json:"club_id"`
		ChessAgeCategory string    `json:"chess_age_category"`

		ELOFideStandard int `json:"elo_fide_standard"`
		ELOFideRapid    int `json:"elo_fide_rapid"`
		ELOFideBlitz    int `json:"elo_fide_blitz"`
		ELOFideBullet   int `json:"elo_fide_bullet"`

		ELONationalStandard int `json:"elo_national_standard"`
		ELONationalRapid    int `json:"elo_national_rapid"`
		ELONationalBlitz    int `json:"elo_national_blitz"`
		ELONationalBullet   int `json:"elo_national_bullet"`

		ELORegionalStandard int    `json:"elo_regional_standard"`
		ELORegionalRapid    int    `json:"elo_regional_rapid"`
		ELORegionalBlitz    int    `json:"elo_regional_blitz"`
		ELORegionalBullet   int    `json:"elo_regional_bullet"`
		Email               string `json:"email"`
		Password            string `json:"-"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), input)
		return
	}
	cnu := data.User{
		FirstName:           input.FirstName,
		LastName:            input.LastName,
		DateOfBirth:         input.DateOfBirth,
		Sex:                 input.Sex,
		ClubID:              input.ClubID,
		ChessAgeCategory:    input.ChessAgeCategory,
		ELOFideStandard:     input.ELOFideStandard,
		ELOFideRapid:        input.ELOFideRapid,
		ELOFideBlitz:        input.ELOFideBlitz,
		ELOFideBullet:       input.ELOFideBullet,
		ELONationalStandard: input.ELONationalStandard,
		ELONationalRapid:    input.ELONationalRapid,
		ELONationalBlitz:    input.ELONationalBlitz,
		ELONationalBullet:   input.ELONationalBullet,
		ELORegionalStandard: input.ELORegionalStandard,
		ELORegionalRapid:    input.ELORegionalRapid,
		ELORegionalBlitz:    input.ELORegionalBlitz,
		ELORegionalBullet:   input.ELORegionalBullet,
		Email:               input.Email,
	}

	// normalize user data before inserting into the database
	normalizeUser(&cnu)

	// create user's unique code
	cnu.UserCode = uuid.New().String()
	// cnu.Password.plainText = cnu.Password

	err := cnu.Password.Set(input.Password)
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), cnu)
		return
	}

	err = app.models.Users.Insert(&cnu)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			apiResponse(c, http.StatusBadRequest, "error", "email exists", cnu)
		default:
			apiResponse(c, http.StatusInternalServerError, "error", err.Error(), cnu)
		}
		return
	}
	// launch go routine to send welcome email in the background
	// this lowers latency of the API response
	app.background(func() {
		// TODO: the welcome template will be based on users' preferred language
		// need to modify user model/databse to include language preference
		err = app.mailer.Send(cnu.Email, "user_welcome_en.tmpl", cnu)
		if err != nil {
			// TODO log error
			// app.logger.Error(err.Error())
		}
	})

	apiResponse(c, http.StatusAccepted, "success", "user created", cnu)
}

func (app *application) getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	email = strings.ToLower(email)
	var user data.User
	err := app.models.Users.GetByEmail(&user)
	if err != nil {
		apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user found", user)
}

func (app *application) updateUser(c *gin.Context) {
	var incommingData map[string]interface{}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	json.Unmarshal(jsonData, &incommingData)
	if _, ok := incommingData["user_code"]; !ok {
		apiResponse(c, http.StatusBadRequest, "error", "user_code is required", incommingData)
		return
	}

	_, err = uuid.Parse(incommingData["user_code"].(string))
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", "invalid user_code", incommingData)
		return
	}

	err = app.models.Users.Update(incommingData)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user updated", incommingData)
}

func (app *application) deleteUser(c *gin.Context) {
	email := c.Param("email")
	email = strings.ToLower(email)
	err := app.models.Users.Delete(email)
	if err != nil {
		apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user deleted", nil)
}
