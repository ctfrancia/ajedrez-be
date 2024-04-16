package main

import (
	"context"
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/mailer"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
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
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config config
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

const version = "1.0.0"

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("CHESS_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")
	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "7005680924ace2", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "f382c81ae48622", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "MCT <no-reply@mct.es>", "SMTP sender")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})
	flag.Parse()

	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(cfg)
	if err != nil {
		// logger.Error("cannot connect to database", "error", err)
		os.Exit(1)
	}

	println("Connected to database")

	defer db.Close()
	app := &application{
		config: cfg,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
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

// TODO: move this to the server.go file
func (app *application) serve() error {
	r := gin.Default()

	r.Use(app.rateLimit())
	r.Use(app.authenticate())
	r.Use(app.enableCORS())
	v1U := r.Group("/v1/user")
	v1T := r.Group("/v1/tournament")
	v1C := r.Group("/v1/club")
	v1Tokens := r.Group("/v1/tokens")
	v1Sys := r.Group("/v1/system")

	// User routes
	// v1U.GET("/all", app.getAllUsers)
	v1U.POST("/create", app.createNewUser)
	v1U.GET("/:email", app.getUserByEmail)
	v1U.PUT("/update/", app.updateUser)
	v1U.DELETE("/delete/:email", app.deleteUser)
	v1U.PUT("/activated", app.activateUser)

	// Tournament routes
	v1T.POST("/create", createNewTournament)

	// Club routes
	// TODO: the middleware below is just for POC, it should be removed
	v1C.Use(app.requireActivatedUser())
	v1C.POST("/create", app.createNewClub)
	v1C.GET("/by-name/:name", app.getClubByName)
	// v1C.GET("/by-code/:code", app.getClubByCode)

	// Token routes
	v1Tokens.POST("/authentication", app.createAuthenticationToken)

	// System routes
	v1Sys.GET("/healthcheck", app.healthcheck)

	port := fmt.Sprintf(":%d", app.config.port)
	// srv.ListenAndServe()
	r.Run(port)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		fmt.Printf("shutting down server, signal: %s", s.String())

		// app.logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Call Shutdown() on the server like before, but now we only send on the
		// shutdownError channel if it returns an error.
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		fmt.Printf("completing background tasks: %s", srv.Addr)

		// Call Wait() to block until our WaitGroup counter is zero --- essentially
		// blocking until the background goroutines have finished. Then we return nil on
		// the shutdownError channel, to indicate that the shutdown completed without
		// any issues.
		app.wg.Wait()
		shutdownError <- nil
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}
