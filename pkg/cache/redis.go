package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//client for redis
type RedisClient struct {
	Client *redis.Client
}

//new redis cache client
func NewRedisClient(ctx context.Context, addr string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisClient{
		Client: client,
	}, nil
}
