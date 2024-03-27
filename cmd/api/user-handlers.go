package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	// "errors"
	"io"
	"strings"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/roserocket/gopartial"
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
	// var oldData *data.User
	var incommingData *data.User

	/*
		var newData *data.User
		if err := c.ShouldBindJSON(&newData); err != nil {
			apiResponse(c, http.StatusBadRequest, "error", err.Error(), newData)
			return
		}
	*/
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	fmt.Println("jsonData", string(jsonData))
	json.Unmarshal(jsonData, &incommingData)
	fmt.Printf("incommingData %#v\n", incommingData)
	// user, err := updateUserPartially(oldData, jsonData)
	err = app.models.Users.Update(incommingData)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		// apiResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	/*
		    if err != nil {
		        apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		        return
		    }
		    fmt.Println("user", user)

			oldData, err := app.models.Users.GetByUserCode(newData.UserCode)
			if err != nil {
				apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
				return
			}
	*/
	/*
		var newData *data.User
		if err := c.ShouldBindJSON(&newData); err != nil {
			apiResponse(c, http.StatusBadRequest, "error", err.Error(), newData)
			return
		}
		_, err := uuid.Parse(newData.UserCode)
		if err != nil {
			apiResponse(c, http.StatusBadRequest, "error", "invalid user code", newData)
			return
		}

		oldData, err = app.models.Users.GetByUserCode(newData.UserCode)
		if err != nil {
			apiResponse(c, http.StatusNotFound, "error", "user not found", nil)
			return
		}

		// u := prepareUserUpdate(oldData, newData)
		// fmt.Println("newData", u)

		err = app.models.Users.Update(oldData, newData)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				apiResponse(c, http.StatusNotFound, "error", "record not found", newData)
				return
			}
			apiResponse(c, http.StatusInternalServerError, "error", err.Error(), newData)
			return
		}

		apiResponse(c, http.StatusCreated, "success", "user updated", newData)
	*/
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
func updateUserPartially(user *data.User, partialDataJSON json.RawMessage) (*data.User, error) {
	var partialData map[string]interface{}
	if err := json.Unmarshal(partialDataJSON, &partialData); err != nil {
		fmt.Println("Error unmarshalling partial data: ", err)
	}
	updatedFields, err := gopartial.PartialUpdate(user, partialData, "json", gopartial.SkipConditions, gopartial.Updaters)
	log.Println("Updated fields: ", updatedFields)

	return user, err
}
