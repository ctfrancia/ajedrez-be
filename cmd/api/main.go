package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/mailer"
	"ctfrancia/ajedrez-be/internal/models"
	"ctfrancia/ajedrez-be/internal/repository"
	"ctfrancia/ajedrez-be/internal/vcs"
	"flag"
	"fmt"
	"log"
	"os"

	"strings"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
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
	config     config
	models     data.Models
	repository repository.Repository
	mailer     mailer.Mailer
	wg         sync.WaitGroup
}

var (
	version = vcs.Version()
)

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
	// IMPORTANT: below should be set to true in production
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", false, "Enable rate limiter")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})
	displayVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	db, err := openDB(cfg)
	if err != nil {
		panic(err)
	}
	println("Connected to database")

	app := &application{
		config: cfg,
		// models: data.NewModels(db),
		repository: repository.NewRepository(db),
		mailer:     mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(cfg config) (*gorm.DB, error) {
	dsn := cfg.db.dsn
	// dsn := "host=localhost user=chess_admin password=pa55word dbname=my_chess_website sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Token{}, &models.Tournament{}, &models.Team{})
	return db, nil
}
