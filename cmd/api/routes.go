package main

import (
	"github.com/gin-contrib/expvar"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()

	r.Use(app.metrics())
	r.Use(app.enableCORS())
	r.Use(app.rateLimit())
	r.Use(app.authenticate())
	v1U := r.Group("/v1/user")
	v1T := r.Group("/v1/tournament")
	v1C := r.Group("/v1/club")
	v1Tokens := r.Group("/v1/tokens")
	v1Sys := r.Group("/v1/system")
	v1Pw := r.Group("/v1/password")

	// User routes
	// v1U.GET("/all", app.getAllUsers)
	v1U.POST("/create", app.createNewUser)
	v1U.GET("/:email", app.getUserByEmail)
	v1U.PUT("/update", app.updateUser)
	v1U.DELETE("/delete/:email", app.deleteUser)
	v1U.PUT("/activated", app.activateUser)

	// Password routes
	v1Pw.POST("/check", app.pwCheck)

	// Tournament routes
	v1T.POST("/create", app.createNewTournament)
	v1T.Use(app.tournamentValidator())
	v1T.PUT("/update", app.updateTournament)

	// Club routes
	// TODO: the middleware below is just for POC, it should be removed
	// v1C.Use(app.requireActivatedUser())
	v1C.POST("/create", app.createNewClub)
	v1C.GET("/by-name/:name", app.getClubByName)
	// v1C.GET("/by-code/:code", app.getClubByCode)

	// Token routes
	v1Tokens.POST("/authentication", app.createAuthenticationToken)

	// System routes
	v1Sys.GET("/healthcheck", app.healthcheck)

	// v1Sys.GET("/version", app.version)
	v1Sys.GET("/debug/vars", expvar.Handler())

	return r
}
