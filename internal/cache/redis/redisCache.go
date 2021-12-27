package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/cookienyancloud/avito-backend-test/pkg/cache"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Cache struct {
	rd *cache.RedisClient
}

func NewCache(rd *cache.RedisClient, ) ICache {
	return &Cache{
		rd: rd,
	}
}

type ICache interface {
	CacheKey(ctx context.Context, key uuid.UUID) error
	CheckKey(ctx context.Context, key uuid.UUID) (bool, error)
}

func (c Cache) CacheKey(ctx context.Context, key uuid.UUID) error {
	fmt.Println("CacheKey", key)
	err := c.rd.Client.Set(ctx, key.String(), true, time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c Cache) CheckKey(ctx context.Context, key uuid.UUID) (bool, error) {
	fmt.Println("CheckKey", key.String())
	var state bool
	err := c.rd.Client.Get(ctx, key.String()).Scan(&state)
	if err != redis.Nil && err != nil {
		fmt.Println("err")
		return false, err
	}
	fmt.Println("state", state)
	return state, nil
}
