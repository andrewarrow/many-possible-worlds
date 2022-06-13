package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var utc *time.Location

func init() {
	utc, _ = time.LoadLocation("UTC")
}

func InsertItem(v *Video, subs, slug string) {

	subsInt, _ := strconv.ParseInt(subs, 10, 64)
	// score is number of subs
	// member is channel_id|channel_title
	subzset := fmt.Sprintf("%s-s", slug)

	// score is published_at int64
	// member is channel_id
	pubzset := fmt.Sprintf("%s-p", slug)

	// score is published_at int64
	// member is video_id
	vidzset := fmt.Sprintf("%s-v", v.ChannelId)

	rz1 := redis.Z{
		Score:  float64(subsInt),
		Member: v.ChannelId + "|" + v.ChannelTitle,
	}
	rz2 := redis.Z{
		Score:  float64(v.PublishedAt),
		Member: v.ChannelId,
	}
	rz3 := redis.Z{
		Score:  float64(v.PublishedAt),
		Member: v.Id,
	}

	nc().ZAdd(ctx, subzset, &rz1).Err()
	nc().ZAdd(ctx, pubzset, &rz2).Err()
	nc().ZAdd(ctx, vidzset, &rz3).Err()

	nc().HSet(ctx, v.Id, "title", v.Title).Err()
	//nc().HSet(ctx, v.Id, "view_count", v.ViewCount).Err()
	//nc().HSet(ctx, v.Id, "c_id", v.ChannelId).Err()

	nc().HSet(ctx, v.ChannelId, "title", v.ChannelTitle).Err()
	//nc().HSet(ctx, v.ChannelId, "subs", subs).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, v.Id, expireTime)
	nc().ExpireAt(ctx, v.ChannelId, expireTime)
}
