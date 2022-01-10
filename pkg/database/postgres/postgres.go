package postgres

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/cookienyancloud/avito-backend-test/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

//postgres database client
func NewClient(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {

	db, err := pgxpool.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(ctx); err != nil {
		return nil, err
	}
	path := filepath.Join("schema", "000001_init_schema.up.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return nil, err
	}
	sql := string(c)
	if _, err := db.Exec(ctx, sql); err != nil {
		return nil, err
	}
	return db, nil
}
