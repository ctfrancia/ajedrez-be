package main

import (
	// "ctfrancia/ajedrez-be/internal/data"
	// "errors"
	"net/http"
	// "strings"
	"github.com/gin-gonic/gin"
)

func (app *application) authenticate(next http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement the authentication middleware
	}
	/*
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// "Vary: Authorization" header in the response indicates to any
			// caches that the response may vary based on the value of the Authorization
			// header in the request.
			w.Header().Add("Vary", "Authorization")

			// returns "" if there is no such header found.
			authorizationHeader := r.Header.Get("Authorization")

			// If no Authorization header is present, we assume that the request is
			// being made by an unauthenticated user. In this case, we call the contextSetUser()
			if authorizationHeader == "" {
				r = app.contextSetUser(r, data.AnonymousUser)
				next.ServeHTTP(w, r)
				return
			}

			// Otherwise, we expect the value of the Authorization header to be in the format
			// "Bearer <token>". We try to split this into its constituent parts, and if the
			// header isn't in the expected format we return a 401 Unauthorized response
			// using the invalidAuthenticationTokenResponse() helper (which we will create
			// in a moment).
			headerParts := strings.Split(authorizationHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				// app.invalidAuthenticationTokenResponse(w, r)
				return
			}

			// Extract the actual authentication token from the header parts.
			token := headerParts[1]

			// Validate the token to make sure it is in a sensible format.
			// v := validator.New()

			// If the token isn't valid, use the invalidAuthenticationTokenResponse()
			// helper to send a response, rather than the failedValidationResponse() helper
			// that we'd normally use.
			if data.ValidateTokenPlaintext(token); err != nil {
				// app.invalidAuthenticationTokenResponse(w, r)
				return
			}

			// Retrieve the details of the user associated with the authentication token,
			// again calling the invalidAuthenticationTokenResponse() helper if no
			// matching record was found. IMPORTANT: Notice that we are using
			// ScopeAuthentication as the first parameter here.
			user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)
			if err != nil {
				switch {
				case errors.Is(err, data.ErrRecordNotFound):
					// japp.invalidAuthenticationTokenResponse(w, r)
				default:
					// app.serverErrorResponse(w, r, err)
				}
				return
			}

			// Call the contextSetUser() helper to add the user information to the request
			// context.
			r = app.contextSetUser(r, user)

			// Call the next handler in the chain.
			next.ServeHTTP(w, r)
		})
	*/
}
