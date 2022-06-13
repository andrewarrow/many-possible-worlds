package redis

import (
	"fmt"
	"time"
)

func QueryDay(slug string) []Video {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("%s-%s", slug, BucketForDay(t))
	return QueryBucket(bucket)
}

func QueryBucket(b string) []Video {
	list := []Video{}

	/*
		members, err := nc().ZRevRangeByScore(ctx, b, &redis.ZRangeBy{
			Min: "-inf",
			Max: "+inf",
		}).Result()

	*/

	vals, err := nc().ZRevRangeWithScores(ctx, b, 0, 100).Result()

	if err != nil {
		fmt.Println(err)
		return list
	}
	for _, item := range vals {
		v := Video{}
		v.Id = item.Member.(string)
		v.PublishedAt = int64(item.Score)
		m := QueryAttributes(v.Id)
		v.Title = m["title"]
		v.ViewCount = m["view_count"]
		v.ChannelId = m["c_id"]
		cmap := QueryAttributes(v.ChannelId)
		v.ChannelTitle = cmap["title"]
		v.Subs = cmap["subs"]
		list = append(list, v)
	}

	return list
}

func QueryAttributes(b string) map[string]string {
	m, err := nc().HGetAll(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return m
	}
	return m
}
