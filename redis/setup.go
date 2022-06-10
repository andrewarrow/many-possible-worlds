package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func nc() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	return rdb
}
