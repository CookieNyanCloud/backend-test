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
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	initEnvErr           = "error initializing env: %v\n"
	initDbErr            = "error initializing database: %v\n"
	serverErr            = "error during http server work: %s\n"
	stopServerErr        = "error trying to stop server: %v\n"
	closeDbConnectionErr = "error closing database connection: %v\n"

	start = "start"
)

func Run(configPath string, local bool) {

	//get env vars from .env
	cfg, err := config.Init(configPath, local)
	if err != nil {
		logger.Errorf(initEnvErr, err)
		return
	}

	//init db
	postgresClient, err := postgres.NewClient(cfg.Postgres)
	if err != nil {
		logger.Errorf(initDbErr, err)
		return
	}
	repos := repository.NewFinanceRepo(postgresClient)

	//init services
	financeService := service.NewFinanceService(repos)
	curService := service.CurService{ApiKey: cfg.ApiKey}

	//http
	handlers := delivery.NewHandler(financeService, curService)

	//server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf(serverErr, err.Error())
		}
	}()
	logger.Info(start)

	//quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	if err := srv.Stop(ctx); err != nil {
		logger.Errorf(stopServerErr, err)
	}
	if err := postgresClient.Close(); err != nil {
		logger.Errorf(closeDbConnectionErr, err.Error())
	}

}
