package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

func InsertWorld(q, slug string) {
	zset := "worlds"
	rz := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: slug,
	}

	nc().ZAdd(ctx, zset, &rz).Err()
	nc().HSet(ctx, slug, "q", q).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, slug, expireTime)
}

func QueryWorlds() []World {
	list := []World{}

	zset := "worlds"
	vals, _ := nc().ZRevRangeWithScores(ctx, zset, int64(0), int64(50)).Result()
	for _, item := range vals {
		w := World{}
		w.Slug = item.Member.(string)
		m := QueryAttributes(w.Slug)
		w.Title = m["q"]
		w.Score = int64(item.Score)
		list = append(list, w)
	}

	return list
}
