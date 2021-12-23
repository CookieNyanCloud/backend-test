package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cookienyancloud/avito-backend-test/internal/cache/redis"
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	delivery "github.com/cookienyancloud/avito-backend-test/internal/delivery/http"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/cookienyancloud/avito-backend-test/pkg/cache"
	"github.com/cookienyancloud/avito-backend-test/pkg/database/postgres"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/cookienyancloud/avito-backend-test/pkg/server"
)

const (
	initEnvErr           = "error initializing env: %w"
	initDbErr            = "error initializing database: %w"
	initCacheErr         = "error initializing cache: %w"
	serverErr            = "error during http server work: %w"
	stopServerErr        = "error trying to stop server: %w"
	closeDbConnectionErr = "error closing database connection: %w"

	start = "start"
)

func Run(configPath string, local bool) {
	ctx := context.Background()

	//init config
	cfg, err := config.Init(configPath, local)
	logger.Errorf(initEnvErr, err)

	//init db
	postgresClient, err := postgres.NewClient(cfg.Postgres)
	logger.Errorf(initDbErr, err)
	repos := repository.NewFinanceRepo(postgresClient)

	//init cache
	cacheClient, err := cache.NewRedisClient(cfg.Redis.Addr, ctx)
	logger.Errorf(initCacheErr, err)
	cacheService := redis.NewCache(cacheClient)

	//init services
	financeService := service.NewFinanceService(repos)
	curService := service.CurService{ApiKey: cfg.ApiKey}

	//http
	handlers := delivery.NewHandler(financeService, curService, cacheService)

	//server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf(serverErr, err)
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
	err = srv.Stop(ctx)
	logger.Errorf(stopServerErr, err)
	err = postgresClient.Close()
	logger.Errorf(closeDbConnectionErr, err)

}
