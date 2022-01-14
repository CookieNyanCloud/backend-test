package redis

import (
	"context"
	"time"

	"github.com/cookienyancloud/avito-backend-test/pkg/cache"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Cache struct {
	rd *cache.RedisClient
}

func NewCache(rd *cache.RedisClient) *Cache {
	return &Cache{
		rd: rd,
	}
}

func (c Cache) CacheKey(ctx context.Context, key uuid.UUID) error {
	return c.rd.Client.Set(ctx, key.String(), true, time.Minute).Err()
}

func (c Cache) CheckKey(ctx context.Context, key uuid.UUID) (bool, error) {
	var state bool
	err := c.rd.Client.Get(ctx, key.String()).Scan(&state)
	if err != redis.Nil && err != nil {
		return false, err
	}
	return state, nil
}
