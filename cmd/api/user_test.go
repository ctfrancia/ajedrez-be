package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUser(t *testing.T) {
	tt := []struct {
		name          string
		app           *application
		input         gin.H
		expectedError error
		expected      gin.H
	}{
		{
			name: "createNewUser",
			app: &application{
				config: config{
					env: "test",
				},
			},
			input: gin.H{
				"FirstName": "John",
				"LastName":  "Doe",
				"Email":     "jd@test.com",
				"Language":  "en",
				"Password":  "password",
			},
			expectedError: nil,
			expected: gin.H{
				"status":  "success",
				"message": "user created",
				"data": gin.H{
					"FirstName": "John",
					"LastName":  "Doe",
					"Email":     "",
					"Language":  "en",
					"UserCode":  "",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			// Execute
			// Verify
			// Teardown
		})
	}
}
