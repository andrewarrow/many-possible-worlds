package redis

import "time"

func SetAuth(uuid string) {
	exp := time.Hour * 24 * 30 * 12 * 2
	nc().Set(ctx, "auth", uuid, exp).Err()
}
func GetAuth() string {
	val, _ := nc().Get(ctx, "auth").Result()
	return val
}
