package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Client *redis.Client
}

func NewClient(addr, pass string) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Client{
		Client: rdb,
	}, nil
}
