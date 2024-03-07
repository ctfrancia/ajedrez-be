package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1U := r.Group("/v1/user")
	v1T := r.Group("/v1/tournament")
	v1C := r.Group("/v1/club")

	// User routes
	v1U.POST("/create", createNewUser)
	v1U.GET("/:email", getUserByEmail)
	v1U.GET("/all", getUsers)
	v1U.PUT("/update/:email", updateUser)

	// Tournament routes
	v1T.POST("/create", createNewTournament)

	// Club routes
	v1C.POST("/create", createNewClub)

	r.Run(":8080")
}
