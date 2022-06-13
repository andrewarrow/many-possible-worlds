package redis

import (
	"fmt"
	"time"
)

func UpdateGemCount(slug, id string, val int) {
	t := time.Now().In(utc)

	bucket := fmt.Sprintf("gem-%s-%s", slug, BucketForDay(t))
	if val > 0 {
		nc().ZIncrBy(ctx, bucket, 1.0, id).Err()
	} else {
		nc().ZRem(ctx, bucket, id).Err()
	}
}
