package redis

import (
	"fmt"
	"time"
)

func QueryDay(slug string, offset int) ([]Video, map[string]Channel) {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("%s-%s", slug, BucketForDay(t))
	return QueryBucket(bucket, offset)
}

func QueryDayGems(slug string) ([]Video, map[string]Channel) {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("gem-%s-%s", slug, BucketForMonth(t))
	return QueryBucket(bucket, 0)
}

func QueryDayCount(slug string) int64 {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("%s-%s", slug, BucketForDay(t))
	count, _ := nc().ZCount(ctx, bucket, "-inf", "+inf").Result()
	return count
}

func QueryBucket(b string, offset int) ([]Video, map[string]Channel) {
	list := []Video{}

	vals, err := nc().ZRevRangeWithScores(ctx, b, int64(offset), int64(offset+50)).Result()
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
