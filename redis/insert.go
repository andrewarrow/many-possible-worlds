package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var utc *time.Location

func init() {
	utc, _ = time.LoadLocation("UTC")
}

func InsertItem(ts int64, v *Video, subs, slug string) {

	t := time.Unix(ts, 0)
	t = t.In(utc)

	bucket := fmt.Sprintf("%s-%s", slug, BucketForDay(t))

	rz := redis.Z{
		Score:  float64(v.PublishedAt),
		Member: v.Id,
	}

	err := nc().ZAdd(ctx, bucket, &rz).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	nc().HSet(ctx, v.Id, "title", v.Title).Err()
	nc().HSet(ctx, v.Id, "view_count", v.ViewCount).Err()
	nc().HSet(ctx, v.Id, "c_id", v.ChannelId).Err()

	nc().HSet(ctx, v.ChannelId, "title", v.ChannelTitle).Err()
	nc().HSet(ctx, v.ChannelId, "subs", subs).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, bucket, expireTime)
	nc().ExpireAt(ctx, v.Id, expireTime)
	nc().ExpireAt(ctx, v.ChannelId, expireTime)
}
