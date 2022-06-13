package redis

import (
	"fmt"
	"time"
)

func UpdateGemCount(slug, id string) {
	t := time.Now().In(utc)

	bucket := fmt.Sprintf("gem-%s-%s", slug, BucketForDay(t))
	nc().ZIncrBy(ctx, bucket, 1.0, id).Err()
}
