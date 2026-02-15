package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-fuego/fuego"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/adapters/repo"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/config"
	httpServer "github.com/course-sphere/course-sphere-backend/services/general/internal/transports/http"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
	"github.com/course-sphere/course-sphere-backend/shared/adapters/external"
)

func gracefulShutdown(apiServer *fuego.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}

func main() {
	cfg, _ := env.ParseAs[config.Config]()
	repo := repo.NewInMemory()
	course := usecase.NewCourse(repo)
	authService := external.AuthService{
		BaseUrl: fmt.Sprintf("%s/auth", cfg.ProxyURL),
	}

	server := httpServer.NewServer(&cfg, course, &authService)

	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err := server.Run()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
