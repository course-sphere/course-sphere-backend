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
	server "github.com/course-sphere/course-sphere-backend/services/general/internal/transports/http"
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
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repo.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	paymentClient := external.HTTPPaymentClient{ProxyURL: cfg.ProxyURL}

	course := usecase.Course{
		Repo:          &repo.Course,
		MaterialRepo:  &repo.Material,
		AttemptRepo:   &repo.Attempt,
		PaymentClient: &paymentClient,
	}
	material := usecase.Material{Repo: &repo.Material}
	question := usecase.Question{Repo: &repo.Question}
	attempt := usecase.Attempt{Repo: &repo.Attempt}
	roadmap := usecase.Roadmap{Repo: &repo.Roadmap}

	authClient := external.HTTPAuthClient{ProxyURL: cfg.ProxyURL}
	userClient := external.HTTPUserClient{ProxyURL: cfg.ProxyURL}

	s := server.Server{
		Config: &cfg,

		Course:   course,
		Material: material,
		Question: question,
		Attempt:  attempt,
		Roadmap:  roadmap,

		AuthClient: &authClient,
		UserClient: &userClient,
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
