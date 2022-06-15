package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var globalClient = getClient()
var lastConnect = time.Now().Unix()

func getClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	lastConnect = time.Now().Unix()

	return rdb
}
func nc() *redis.Client {
	if time.Now().Unix()-lastConnect > 3600 {
		globalClient = getClient()
	}
	return globalClient
}
