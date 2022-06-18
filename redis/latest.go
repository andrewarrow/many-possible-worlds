package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Latest struct {
	ChannelId               string
	ChannelTitle            string
	ImageUrl                string
	ExampleVideoId          string
	ExampleVideoPublishedAt int64
	ExampleVideoTitle       string
	SubscriberCount         int64
	ViewCount               int64
	VideoCount              int64
}

func LoadLatest(id string) *Latest {
	l := Latest{}
	l.ChannelId = id
	m := QueryAttributes(l.ChannelId)
	l.ImageUrl = m["img"]
	l.ChannelTitle = m["title"]
	l.ExampleVideoId = m["vid"]
	l.ExampleVideoPublishedAt, _ = strconv.ParseInt(m["vat"], 10, 64)
	l.ExampleVideoTitle = m["vt"]
	l.SubscriberCount, _ = strconv.ParseInt(m["subs"], 10, 64)
	l.ViewCount, _ = strconv.ParseInt(m["vc"], 10, 64)
	l.VideoCount, _ = strconv.ParseInt(m["vidc"], 10, 64)
	return &l
}

func QueryLatest(zset string, amount int) []*Latest {
	list := []*Latest{}

	vals, _ := nc().ZRevRangeWithScores(ctx, zset, int64(0), int64(amount-1)).Result()
	for _, item := range vals {
		id := item.Member.(string)
		l := LoadLatest(id)
		//w.Score = int64(item.Score)
		list = append(list, l)
	}

	return list
}

func UpdateLatestVc(l *Latest) {
	zset := "latest-vc"
	rz := redis.Z{Score: float64(l.ViewCount), Member: l.ChannelId}
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

func InsertLatest(l *Latest) {

	zset := "latest"
	rz := redis.Z{Score: float64(time.Now().Unix()), Member: l.ChannelId}
	nc().ZAdd(ctx, zset, &rz).Err()

	zset = "latest-vc"
	rz = redis.Z{Score: float64(l.ViewCount), Member: l.ChannelId}
	nc().ZAdd(ctx, zset, &rz).Err()

	nc().HSet(ctx, l.ChannelId, "title", l.ChannelTitle).Err()
	nc().HSet(ctx, l.ChannelId, "img", l.ImageUrl).Err()
	nc().HSet(ctx, l.ChannelId, "vid", l.ExampleVideoId).Err()
	nc().HSet(ctx, l.ChannelId, "vat", l.ExampleVideoPublishedAt).Err()
	nc().HSet(ctx, l.ChannelId, "vt", l.ExampleVideoTitle).Err()
	nc().HSet(ctx, l.ChannelId, "subs", fmt.Sprintf("%d", l.SubscriberCount)).Err()
	nc().HSet(ctx, l.ChannelId, "vc", fmt.Sprintf("%d", l.ViewCount)).Err()
	nc().HSet(ctx, l.ChannelId, "vidc", fmt.Sprintf("%d", l.VideoCount)).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	//nc().ExpireAt(ctx, v.Id, expireTime)
	nc().ExpireAt(ctx, l.ChannelId, expireTime)
}
