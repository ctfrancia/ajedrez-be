package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}
type application struct {
	config config
	logger *slog.Logger
}

const version = "1.0.0"

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	// Database connection string: postgres://username:password@localhost:5432/database_name
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://chess:che55@localhost:5432/chess?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	port := fmt.Sprintf(":%d", cfg.port)
	r.Run(port)
}

func openDB(cfg config) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	// Return the sql.DB connection pool.
	return db, nil

}
