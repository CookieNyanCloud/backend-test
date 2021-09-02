package app

import(
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/cookienyancloud/avito-backend-test/pkg/database/postgres"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string)  {
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


}