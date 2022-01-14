package postgresql

import (
	"context"
	"fmt"

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
	return db, nil
}
