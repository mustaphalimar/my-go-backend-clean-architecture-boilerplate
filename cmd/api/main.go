package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/mustaphalimar/prepilotapp-backend/pkg/config"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/handlers"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/repositories"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/router"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/usecases"
)

const DefaultContextTimeout = 30

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	srv, err := server.New(cfg)
	if err != nil {
		log.Fatal("failed to initialize server")
	}

	// Initialize repositories, services, and handlers
	repos := repositories.NewRepositories(srv)
	usecases := usecases.NewUsecases(repos)
	handlers := handlers.NewHandlers(srv, usecases)

	r := router.NewRouter(srv, handlers, usecases)

	// Setup HTTP server
	srv.SetupHTTPServer(r)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// Start server
	go func() {
		if err = srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeout*time.Second)
	defer cancel()
	defer stop()

	log.Println("shutting down server gracefully")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown")
	}

	log.Println("server exited")
}
