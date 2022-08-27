package redis

import (
	"context"
	"time"

	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"

	"github.com/go-redis/redis"
	"go.elastic.co/apm/module/apmgoredis/v2"
)

type Redis struct {
	address  string
	password string
	conn     apmgoredis.Client
}

func New(address, password string) ports.Catch {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return &Redis{
		address:  address,
		password: password,
		conn:     apmgoredis.Wrap(client),
	}
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	c := r.conn.WithContext(ctx)
	res := c.Get(key)
	if res.Err() != nil {
		return "", res.Err()
	}
	return res.Val(), nil
}

func (r *Redis) Set(ctx context.Context, key string, data interface{}, exp time.Duration) error {
	c := r.conn.WithContext(ctx)
	res := c.Set(key, data, exp)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
