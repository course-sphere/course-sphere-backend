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

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/adapters/presign"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/config"
	server "github.com/course-sphere/course-sphere-backend/services/storage/internal/transports/http"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/usecase"
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
	ctx := context.Background()

	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		log.Fatal(err)
	}

	presigner, err := presign.NewS3PresignClient(ctx, cfg.S3Endpoint, cfg.S3Bucket)
	if err != nil {
		log.Fatal(err)
	}
	presign := usecase.Presign{Presigner: presigner}

	s := server.Server{
		Config:  &cfg,
		Presign: presign,
	}
	server := s.Build()

	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err = server.Run()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
