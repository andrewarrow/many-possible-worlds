package redis

import (
	"fmt"
	"time"
)

func QueryDayCount(slug string) int64 {
	t := time.Now().In(utc)
	bucket := fmt.Sprintf("%s-%s", slug, BucketForDay(t))
	count, _ := nc().ZCount(ctx, bucket, "-inf", "+inf").Result()
	return count
}

func QueryPinned(slug string) []Channel {
	list := []Channel{}

	zset := fmt.Sprintf("gem-%s", slug)
	vals, _ := nc().ZRevRangeWithScores(ctx, zset, int64(0), int64(9)).Result()
	for _, item := range vals {
		c := Channel{}
		c.Id = item.Member.(string)
		m := QueryAttributes(c.Id)
		c.Title = m["title"]
		c.ImageUrl = m["img"]
		c.SubscriberCount = m["subs"]
		list = append(list, c)
	}

	return list
}

func QueryChannelsInSlug(slug string, offset int) []Channel {
	list := []Channel{}

	pubzset := fmt.Sprintf("%s-p", slug)
	vals, _ := nc().ZRevRangeWithScores(ctx, pubzset, int64(offset), int64(offset+50)).Result()
	for _, item := range vals {
		c := Channel{}
		c.Id = item.Member.(string)
		m := QueryAttributes(c.Id)
		c.Title = m["title"]
		c.SubscriberCount = m["subs"]
		c.ImageUrl = m["img"]
		c.PublishedAt = int64(item.Score)
		list = append(list, c)
	}

	return list
}

func QueryVideosInChannel(cid string, offset int) []Video {
	list := []Video{}

	vidzset := fmt.Sprintf("%s-v", cid)
	vals, _ := nc().ZRevRangeWithScores(ctx, vidzset, int64(0), int64(50)).Result()
	for _, item := range vals {
		v := Video{}
		v.Id = item.Member.(string)
		v.PublishedAt = int64(item.Score)
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
