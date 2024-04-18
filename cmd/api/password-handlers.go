package main

import (
	"github.com/gin-gonic/gin"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/http"
)

func (app *application) pwCheck(c *gin.Context) {
	const minEntropyBits = 60
	var valid string
	var input struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		apiResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	entropy := passwordvalidator.GetEntropy(input.Password)

	err := passwordvalidator.Validate(input.Password, minEntropyBits)
	if err != nil {
		valid = err.Error()
	} else {
		valid = "password strong enough"
	}

	resp := map[string]interface{}{
		"entropy":   entropy,
		"is_strong": nil == err,
		"notes":     valid,
	}

	apiResponse(c, http.StatusOK, "success", "password check complete", resp)
}
