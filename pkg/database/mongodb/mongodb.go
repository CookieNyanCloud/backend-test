package mongodb

import (
	"context"
	"fmt"

	"github.com/cookienyancloud/avito-backend-test/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, cfg config.MongoConfig) (*mongo.Client, error) {
	var connstr string
	if cfg.Username == "" && cfg.Password == "" {
		connstr = "mongodb://%s:%s"
	} else {
		connstr = "mongodb://%s:%s@%s:%s"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connstr))
	if err != nil {
		return nil, fmt.Errorf("mongo new client: %w", err)
	}
	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping: %w", err)
	}
	return client, nil
}
