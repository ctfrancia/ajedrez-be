package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
