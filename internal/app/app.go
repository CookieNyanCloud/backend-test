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
	delivery "github.com/cookienyancloud/avito-backend-test/internal/delivery/httprest"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/cookienyancloud/avito-backend-test/pkg/cache"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/cookienyancloud/avito-backend-test/pkg/server"
)

//application initiation
func Run(configPath string, local bool) {
	ctx := context.Background()

	//init config
	cfg, err := config.Init(configPath, local)
	logger.Errorf("error initializing env: %v", err)

	//init db
	repo, err := repository.SwitchDb(ctx, cfg)
	logger.Errorf("initializing db: %v", err)

	//init cache
	cacheClient, err := cache.NewRedisClient(ctx, cfg.Redis.Addr)
	logger.Errorf("error initializing cache: %v", err)
	cacheService := redis.NewCache(cacheClient)

	//init services
	financeService := service.NewFinanceService(repo)
	curService := service.NewCurService(cfg.ApiKey)

	//http
	handlers := delivery.NewHandler(financeService, curService, cacheService)

	//server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error during http server work: %v", err)
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
	logger.Errorf("error trying to stop server: %v", err)
	err = repo.Close(ctx)
	logger.Errorf("error closing database connection: %v", err)

}
