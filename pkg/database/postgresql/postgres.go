package postgresql

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/cookienyancloud/avito-backend-test/internal/config"
	"github.com/jmoiron/sqlx"
)

//postgres database client
func NewClient(ctx context.Context, cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	path := filepath.Join("schema", "000001_init_schema.up.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return nil, err
	}
	sql := string(c)
	if _, err := db.ExecContext(ctx, sql); err != nil {
		return nil, err
	}
	return db, nil
}
