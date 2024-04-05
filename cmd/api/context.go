package main

import (
	// "context"
	"ctfrancia/ajedrez-be/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Define a custom contextKey type, with the underlying type string.
type contextKey string

// We'll use this constant as the key for getting and setting user information
// in the request context.
const userContextKey = contextKey("user")

// The contextSetUser() method returns a new copy of the request with the provided
// User struct added to the context. Note that we use our userContextKey constant as the
// key.
func (app *application) contextSetUser(c *gin.Context, user *data.User) *gin.Context {
	// ctx := context.WithValue(c, userContextKey, user)
	return c // ctx
}

func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
