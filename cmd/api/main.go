package main

import (
	"context"
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/mailer"
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}
type application struct {
	config config
	models data.Models
	mailer mailer.Mailer
}

const version = "1.0.0"

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("CHESS_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")
	// Read the SMTP server configuration settings into the config struct, using the
	// Mailtrap settings as the default values. IMPORTANT: If you're following along,
	// make sure to replace the default values for smtp-username and smtp-password
	// with your own Mailtrap credentials.
	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "7005680924ace2", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "f382c81ae48622", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "MCT <no-reply@mct.es>", "SMTP sender")
	flag.Parse()

	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(cfg)
	if err != nil {
		// logger.Error("cannot connect to database", "error", err)
		os.Exit(1)
	}

	defer db.Close()
	app := &application{
		config: cfg,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	r := gin.Default()
	v1U := r.Group("/v1/user")
	v1T := r.Group("/v1/tournament")
	v1C := r.Group("/v1/club")

	// User routes
	// v1U.GET("/all", app.getAllUsers)
	v1U.POST("/create", app.createNewUser)
	v1U.GET("/:email", app.getUserByEmail)
	v1U.PUT("/update/", app.updateUser)
	v1U.DELETE("/delete/:email", app.deleteUser)

	// Tournament routes
	v1T.POST("/create", createNewTournament)

	// Club routes
	v1C.POST("/create", app.createNewClub)
	v1C.GET("/by-name/:name", app.getClubByName)
	// v1C.GET("/by-code/:code", app.getClubByCode)

	port := fmt.Sprintf(":%d", cfg.port)
	r.Run(port)
}

func openDB(cfg config) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Set the maximum idle timeout for connections in the pool. Passing a duration less
	// than or equal to 0 will mean that connections are not closed due to their idle time.
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

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
