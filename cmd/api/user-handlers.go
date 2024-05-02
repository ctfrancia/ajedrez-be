package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/models"
	"ctfrancia/ajedrez-be/internal/repository"
	"ctfrancia/ajedrez-be/pkg/dtos"
	"net/http"

	"io"
	"log"
	"strings"

	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createNewUser(c *gin.Context) {
	var input dtos.UserCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), input)
		return
	}
	cnu := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     &input.Email,
		Language:  input.Language,
		UserCode:  uuid.New().String(),
	}

	// normalize user data before inserting into the database
	normalizeUser(&cnu)

	// err := cnu.Password.Set(input.Password)
	hashed, err := models.PasswordSet(input.Password)
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), input)
		return
	}

	cnu.Password = hashed
	err = app.repository.Users.Create(&cnu)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateEmail):
			resp := map[string]interface{}{
				"email": cnu.Email,
			}
			apiResponse(c, http.StatusBadRequest, "error", "email exists", resp)
		default:
			apiResponse(c, http.StatusInternalServerError, "error", err.Error(), input)
		}

		return
	}

	// After the user record has been created in the database, generate a new activation
	// token for the user.
	token, err := app.repository.Tokens.New(cnu.ID, models.TokenLifetime, data.ScopeActivation)
	if err != nil {
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	app.background(func() {
		data := map[string]any{
			"activationToken": token.Plaintext,
			"userID":          cnu.ID,
		}
		// TODO: the welcome template will be based on users' preferred language
		err = app.mailer.Send(*cnu.Email, "user_welcome_en.tmpl", data)
		if err != nil {
			log.Fatal(err)
			// TODO: send error to monitoring service
			// app.logger.Error(err.Error())
		}
	})

	resp := map[string]interface{}{
		"user_code": cnu.UserCode,
	}
	apiResponse(c, http.StatusCreated, "success", "user created", resp)
}

func (app *application) getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	email = strings.ToLower(email)
	var user *data.User
	user, err := app.models.Users.GetByEmail(email)
	if err != nil {
		apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user found", user)
}

func (app *application) updateUser(c *gin.Context) {
	var input map[string]interface{}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	json.Unmarshal(jsonData, &input)
	if _, ok := input["user_code"]; !ok {
		apiResponse(c, http.StatusBadRequest, "error", "user_code is required", input)
		return
	}

	_, err = uuid.Parse(input["user_code"].(string))
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", "invalid user_code", input)
		return
	}

	err = app.repository.Users.Update(input)
	if err != nil {
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user updated", input)
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

func (app *application) activateUser(c *gin.Context) {
	var input dtos.ActivateTokenDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	if err := data.ValidateTokenPlaintext(input.Token); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	user, err := app.repository.Users.GetForToken(data.ScopeActivation, input.Token)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
		default:
			apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		}
		return
	}

	u := map[string]interface{}{
		"id":        user.ID,
		"activated": true,
	}

	err = app.repository.Users.Update(u)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			apiResponse(c, http.StatusConflict, "error", "edit conflict", nil)

		case errors.Is(err, data.ErrRecordNotFound):
			apiResponse(c, http.StatusNotFound, "error", err.Error(), nil)

		default:
			apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		}
		return
	}

	if app.repository.Tokens.DeleteAllForUser(data.ScopeActivation, user.ID); err != nil {
		apiResponse(c, http.StatusNotFound, "error", err.Error(), nil)
		return
	}

	apiResponse(c, http.StatusOK, "success", "user activated", nil)
}
