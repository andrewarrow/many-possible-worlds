package redis

import (
	"fmt"
	"time"
)

func UpdateGemCount(slug, id string, val int) {
	t := time.Now().In(utc)

	bucket := fmt.Sprintf("gem-%s-%s", slug, BucketForMonth(t))
	if val > 0 {
		nc().ZIncrBy(ctx, bucket, 1.0, id).Err()
	} else {
		nc().ZRem(ctx, bucket, id).Err()
	}

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, bucket, expireTime)
}
