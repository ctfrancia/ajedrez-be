package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/models"
	"ctfrancia/ajedrez-be/pkg/dtos"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) createAuthenticationToken(c *gin.Context) {
	var input dtos.AuthenticateUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	user, err := app.repository.Users.GetByEmail(input.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			app.notFoundResponse(c)
			return

		default:
			app.internalServerError(c, "")
			return
		}
	}

	match, err := models.PasswordMatches(user.Password, input.Password) //user.Password.Matches(input.Password)
	if err != nil {
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	if !match {
		apiResponse(c, http.StatusUnauthorized, "error", data.ErrInvalidCredentials, nil)
		return
	}

	token, err := app.repository.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	response := map[string]string{
		"token":  token.Plaintext,
		"expiry": token.Expiry.Format(time.RFC3339),
	}
	apiResponse(c, http.StatusOK, "success", "", response)
}
