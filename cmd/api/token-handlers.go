package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (app *application) createAuthenticationToken(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
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

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	if !match {
		apiResponse(c, http.StatusUnauthorized, "error", data.ErrInvalidCredentials, nil)
		return
	}

	token, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
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
