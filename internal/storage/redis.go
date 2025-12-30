package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func New(addr string) *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl int64) error {
	return r.Client.Set(ctx, key, value, time.Duration(ttl)*time.Millisecond).Err()
}
