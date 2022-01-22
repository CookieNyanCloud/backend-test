package postgresql

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

//postgres database client
func NewClient(ctx context.Context, username, password, host, port, dbName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			username, password, host, port, dbName))
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
