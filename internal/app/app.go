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

//application initiation
func Run(configPath string, local bool) {
	ctx := context.Background()

	//init config
	cfg, err := config.Init(configPath, local)
	logger.Errorf("error initializing env: %w", err)

	//init db
	postgresClient, err := postgres.NewClient(cfg.Postgres)
	logger.Errorf("error initializing database: %w", err)
	repos := repository.NewFinanceRepo(postgresClient)

	//init cache
	cacheClient, err := cache.NewRedisClient(ctx, cfg.Redis.Addr)
	logger.Errorf("error initializing cache: %w", err)
	cacheService := redis.NewCache(cacheClient)

	//init services
	financeService := service.NewFinanceService(repos)
	curService := service.NewCurService(cfg.ApiKey)

	//http
	handlers := delivery.NewHandler(financeService, curService, cacheService)

	//server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error during http server work: %w", err)
		}
	}()
	logger.Info("start")

	//quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Info("stop")
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	err = srv.Stop(ctx)
	logger.Errorf("error trying to stop server: %w", err)
	err = postgresClient.Close()
	logger.Errorf("error closing database connection: %w", err)

}
