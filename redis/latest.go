package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func QueryLatestVideos(zset string, amount int) []*Video {
	list := []*Video{}

	vals, _ := nc().ZRevRangeWithScores(ctx, zset, int64(0), int64(amount-1)).Result()
	for _, item := range vals {
		id := item.Member.(string)
		v := LoadVideo(id)
		//w.Score = int64(item.Score)
		list = append(list, v)
	}

	return list
}

func QueryLatest(zset string, amount int) []*Channel {
	list := []*Channel{}

	vals, _ := nc().ZRevRangeWithScores(ctx, zset, int64(0), int64(amount-1)).Result()
	for _, item := range vals {
		id := item.Member.(string)
		c := LoadChannel(id)
		//w.Score = int64(item.Score)
		list = append(list, c)
	}

	return list
}

func UpdateLatestVc(c *Channel) {
	zset := "latest-vc"
	rz := redis.Z{Score: float64(c.ViewCount), Member: c.Id}
	nc().ZAdd(ctx, zset, &rz).Err()
}

func UpdateLatest(id string, viewCount int64) {
	if id == "" {
		return
	}
	if viewCount > 99999 {
		return
	}
	zset := "latest"
	rz := redis.Z{Score: float64(time.Now().Unix()), Member: id}
	nc().ZAdd(ctx, zset, &rz).Err()
}

func InsertLatest(c *Channel) {

	zset := "latest"
	rz := redis.Z{Score: float64(time.Now().Unix()), Member: c.Id}
	nc().ZAdd(ctx, zset, &rz).Err()

	zset = "latest-vc"
	rz = redis.Z{Score: float64(c.ViewCount), Member: c.Id}
	nc().ZAdd(ctx, zset, &rz).Err()

	nc().HSet(ctx, c.Id, "title", c.Title).Err()
	nc().HSet(ctx, c.Id, "img", c.ImageUrl).Err()
	nc().HSet(ctx, c.Id, "vid", c.ExampleVideoId).Err()
	nc().HSet(ctx, c.Id, "vat", c.ExampleVideoPublishedAt).Err()
	nc().HSet(ctx, c.Id, "vt", c.ExampleVideoTitle).Err()
	nc().HSet(ctx, c.Id, "subs", fmt.Sprintf("%d", c.SubscriberCount)).Err()
	nc().HSet(ctx, c.Id, "vc", fmt.Sprintf("%d", c.ViewCount)).Err()
	nc().HSet(ctx, c.Id, "vidc", fmt.Sprintf("%d", c.VideoCount)).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	//nc().ExpireAt(ctx, v.Id, expireTime)
	nc().ExpireAt(ctx, c.Id, expireTime)
}
