package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"fmt"
	"net/http"

	"io"
	"strings"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createNewUser(c *gin.Context) {
	var cnu data.User
	if err := c.ShouldBindJSON(&cnu); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), cnu)
		return
	}

	// normalize user data before inserting into the database
	normalizeUser(&cnu)

	// create user's unique code
	cnu.UserCode = uuid.New().String()

	err := app.models.Users.Insert(&cnu)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_club_user_code_unique\"" {
			apiResponse(c, http.StatusBadRequest, "error", "user already exists", cnu)
			return
		}
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), cnu)
		return
	}

	apiResponse(c, http.StatusCreated, "success", "user created", cnu)
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

	vn, err := app.models.Users.Update(incommingData)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	incommingData["version"] = vn
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
