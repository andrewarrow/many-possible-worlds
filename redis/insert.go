package redis

import (
	"fmt"
	"time"
)

var utc *time.Location

func init() {
	utc, _ = time.LoadLocation("UTC")
}

func InsertItem(ts int64, videoId, title, channelTitle, viewCount, channelId, subs string) {

	t := time.Unix(ts, 0)
	t = t.In(utc)

	bucket := BucketForHour(t)

	err := nc().SAdd(ctx, bucket, videoId).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	nc().HSet(ctx, videoId, "title", title).Err()
	nc().HSet(ctx, videoId, "view_count", viewCount).Err()
	nc().HSet(ctx, videoId, "c_id", channelId).Err()

	nc().HSet(ctx, channelId, "title", channelTitle).Err()
	nc().HSet(ctx, channelId, "subs", subs).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, bucket, expireTime)
	nc().ExpireAt(ctx, videoId, expireTime)
	nc().ExpireAt(ctx, channelId, expireTime)

}
