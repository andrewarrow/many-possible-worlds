package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func LoadVideo(id string) *Video {
	v := Video{}
	m := QueryAttributes(id)
	v.ImageUrl = m["img"]
	v.Title = m["title"]
	v.ChannelId = m["cid"]
	v.PublishedAt, _ = strconv.ParseInt(m["pub"], 10, 64)
	v.Id = id
	return &v
}

func StoreSingleVideo(v *Video) {
	nc().HSet(ctx, v.Id, "title", v.Title).Err()
	nc().HSet(ctx, v.Id, "cid", v.ChannelId).Err()
	nc().HSet(ctx, v.Id, "img", v.ImageUrl).Err()
	nc().HSet(ctx, v.Id, "pub", v.PublishedAt).Err()
	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, v.Id, expireTime)

	AddToVideoSet(v.ChannelId, v.Id, v.PublishedAt)
}

func AddToVideoSet(cid, vid string, pub int64) {
	vidzset := fmt.Sprintf("%s-v", cid)
	rz := redis.Z{Score: float64(pub), Member: vid}
	nc().ZAdd(ctx, vidzset, &rz).Err()
}

func FindOlderVideo(cid string, pub int64) *Video {
	vidzset := fmt.Sprintf("%s-v", cid)
	zrb := redis.ZRangeBy{
		Max:    fmt.Sprintf("%d", pub),
		Min:    "-inf",
		Count:  1,
		Offset: 0,
	}
	vals, _ := nc().ZRangeByScore(ctx, vidzset, &zrb).Result()
	if len(vals) == 0 {
		return nil
	}
	return LoadVideo(vals[0])
}
