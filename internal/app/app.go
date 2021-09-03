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

func Run(configPath string) {

	//подтягиваем значения переменных из папки конфигураций и .env
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Errorf("ошибка инициализации переменных:%v",err)
		return
	}
	println(cfg.Postgres.SSLMode)
	println(cfg.Postgres.Password)
	println(cfg.Postgres.Username)
	println(cfg.Postgres.DBName)
	println(cfg.Postgres.Host)
	println(cfg.Postgres.Port)

	//инициализация базы данных
	postgresClient, err := postgres.NewClient(cfg.Postgres)
	if err != nil {
		logger.Errorf("ошибка инициализации базы данных:%v",err)
		return
	}
	repos := repository.NewFinanceRepo(postgresClient)

	//инициализация сервиса
	financeService := service.NewFinanceService(repos)

	//http
	handlers := delivery.NewHandler(financeService)

	//сервер
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("возникла ошибка в работе http сервера: %s\n", err.Error())
		}
	}()
	logger.Info("запуск")

	//выход
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("ошибка при остановке сервера: %v",err)
	}
	if err := postgresClient.Close(); err != nil {
		logger.Errorf("ошибка при закрытии соединения с бд: %v",err.Error())
	}

}
