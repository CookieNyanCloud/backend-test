package postgres

import (
	"fmt"
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"path/filepath"
)

func NewClient(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.DBName,
			cfg.Password,
			cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	path := filepath.Join("schema", "000001_init_schema.up.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return nil, err
	}
	sql := string(c)
	sqlx.MustExec(db, sql)

	return db, nil
}
