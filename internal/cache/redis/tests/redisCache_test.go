package tests

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/cookienyancloud/avito-backend-test/internal/cache/redis"
	"github.com/cookienyancloud/avito-backend-test/pkg/cache"
	"github.com/go-redis/redismock/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCacheKey(t *testing.T) {
	type mockBehavior func(key uuid.UUID, val bool, time time.Duration)
	db, mock := redismock.NewClientMock()
	tt := []struct {
		name   string
		key    uuid.UUID
		val    bool
		dur    time.Duration
		mockB  mockBehavior
		expErr bool
	}{
		{
			name: "ok",
			key:  uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
			val:  true,
			dur:  time.Minute,
			mockB: func(key uuid.UUID, val bool, time time.Duration) {
				mock.
					ExpectSet(key.String(), val, time).RedisNil()
			},
			expErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := &cache.RedisClient{Client: db}
			mockRedis := redis.NewCache(mockClient)
			tc.mockB(tc.key, tc.val, tc.dur)
			err := mockRedis.CacheKey(context.Background(), tc.key)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckKey(t *testing.T) {
	type mockBehavior func(key uuid.UUID)
	db, mock := redismock.NewClientMock()
	tt := []struct {
		name     string
		key      uuid.UUID
		mockB    mockBehavior
		expState string
		expErr   bool
	}{
		{
			name: "nil",
			key:  uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
			mockB: func(key uuid.UUID) {
				mock.
					ExpectGet(key.String()).RedisNil()
			},
			expState: "false",
			expErr:   false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := &cache.RedisClient{Client: db}
			mockRedis := redis.NewCache(mockClient)
			tc.mockB(tc.key)
			state, err := mockRedis.CheckKey(context.Background(), tc.key)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
			boolState := strconv.FormatBool(state)
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, boolState, tc.expState)
			}
		})
	}
}
