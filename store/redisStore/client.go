package redisStore

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient(addr, pass string) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &RedisClient{
		Client: rdb,
		Ctx:    ctx,
	}, nil
}
