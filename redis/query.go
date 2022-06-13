package redis

import (
	"fmt"
	"time"
)

func QueryDay(slug string) ([]Video, map[string]Channel) {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("%s-%s", slug, BucketForDay(t))
	return QueryBucket(bucket)
}

func QueryDayGems(slug string) ([]Video, map[string]Channel) {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("gem-%s-%s", slug, BucketForDay(t))
	return QueryBucket(bucket)
}

func QueryBucket(b string) ([]Video, map[string]Channel) {
	list := []Video{}

	/*
		members, err := nc().ZRevRangeByScore(ctx, b, &redis.ZRangeBy{
			Min: "-inf",
			Max: "+inf",
		}).Result()

	*/

	vals, err := nc().ZRevRangeWithScores(ctx, b, 0, 100).Result()
	cmap := map[string]Channel{}

	if err != nil {
		fmt.Println(err)
		return list, cmap
	}
	cidmap := map[string]bool{}
	for _, item := range vals {
		v := Video{}
		v.Id = item.Member.(string)
		v.PublishedAt = int64(item.Score)
		m := QueryAttributes(v.Id)
		v.Title = m["title"]
		v.ViewCount = m["view_count"]
		v.ChannelId = m["c_id"]
		cidmap[v.ChannelId] = true
		list = append(list, v)
	}

	for cid, _ := range cidmap {
		m := QueryAttributes(cid)
		c := Channel{}
		c.Title = m["title"]
		c.SubscriberCount = m["subs"]
		cmap[cid] = c
	}

	return list, cmap
}

func QueryAttributes(b string) map[string]string {
	m, err := nc().HGetAll(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return m
	}
	return m
}
