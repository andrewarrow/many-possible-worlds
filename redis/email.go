package redis

import "time"

func LookupEmail(email string) string {
	m, _ := nc().HGetAll(ctx, email).Result()
	if len(m) == 0 {
		return ""
	}
	return m["pass"]
}

func SaveEmailPassword(email, password string) {
	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().HSet(ctx, email, "pass", password).Err()
	nc().ExpireAt(ctx, email, expireTime)
}
