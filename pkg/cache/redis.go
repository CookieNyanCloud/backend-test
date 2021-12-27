package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	Client *redis.Client
}

//new redis cache client
func NewRedisClient(ctx context.Context, addr string) (*redisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &redisClient{
		Client: client,
	}, nil
}
