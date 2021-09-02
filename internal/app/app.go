package app

import (
	"context"
	"errors"
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	delivery "github.com/cookienyancloud/avito-backend-test/internal/delivery/http"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/cookienyancloud/avito-backend-test/internal/server"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/cookienyancloud/avito-backend-test/pkg/database/postgres"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	//log.SetFormatter(&log.JSONFormatter{})

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Error(err)
		return
	}

	postgresClient, err := postgres.NewClient(cfg.Postgres)
	if err != nil {
		log.Error(err)
		return
	}

	repos := repository.NewRepositories(postgresClient)

	service := service.NewUsersService(repos)
	handlers := delivery.NewHandler(service)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}

	if err := postgresClient.Close(); err != nil {
		log.Errorf("error occurred on db connection close: %s", err.Error())
	}

}
