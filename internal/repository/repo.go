package repository

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/config"
	"github.com/cookienyancloud/avito-backend-test/internal/repository/mongo"
	"github.com/cookienyancloud/avito-backend-test/internal/repository/postgres"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/cookienyancloud/avito-backend-test/pkg/database/mongodb"
	"github.com/cookienyancloud/avito-backend-test/pkg/database/postgresql"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/pkg/errors"
)

func SwitchDb(ctx context.Context, cfg *config.Config) (service.IRepo, error) {
	var repo service.IRepo
	switch cfg.State.DataBase {
	case "postgres":
		logger.Info("postgres")
		postgresClient, err := postgresql.NewClient(ctx, cfg.Postgres)
		if err != nil {
			return nil, errors.Wrap(err, "postgres")
		}
		repo = postgres.NewFinanceRepo(postgresClient)
	case "mongo":
		logger.Info("mongo")
		mongoClient, err := mongodb.NewClient(ctx, cfg.Mongo)
		if err != nil {
			return nil, errors.Wrap(err, "mongo")
		}
		repo = mongo.NewFinanceRepo(mongoClient)
	default:
		return nil, errors.New("no such database")
	}

	return repo, nil
}
