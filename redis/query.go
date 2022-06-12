package redis

import (
	"fmt"
	"time"
)

func QueryHour() []Video {
	t := time.Now().In(utc)
	t = t.Add(time.Hour * -1)
	bucket := BucketForHour(t)
	return QueryBucket(bucket)
}

func QueryBucket(b string) []Video {
	list := []Video{}
	members, err := nc().SMembers(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return list
	}
	for _, item := range members {
		v := Video{}
		v.Id = item
		m := QueryAttributes(v.Id)
		v.Title = m["title"]
		v.ViewCount = m["view_count"]
		v.ChannelId = m["c_id"]
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
