package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func UpdateGemCount(slug, id string, val int) {
	bucket := fmt.Sprintf("gem-%s", slug)
	if val > 0 {
		rz := redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: id,
		}

		nc().ZAdd(ctx, bucket, &rz).Err()
	} else {
		nc().ZRem(ctx, bucket, id).Err()
	}

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, bucket, expireTime)
}

func UpdateChannelImage(id, url string) {
	nc().HSet(ctx, id, "img", url).Err()
}
