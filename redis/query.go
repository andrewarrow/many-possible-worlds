package redis

import (
	"fmt"
	"time"
)

type RedisVideo struct {
	Id string
}

func QueryHour() []RedisVideo {
	t := time.Now().In(utc)
	t = t.Add(time.Hour * -1)
	bucket := BucketForHour(t)
	items := QueryBucket(bucket)

	list := []RedisVideo{}
	for _, item := range items {
		v := RedisVideo{}
		v.Id = item
		list = append(list, v)
	}

	return list
}

func QueryBucket(b string) []string {
	items := []string{}
	members, err := nc().SMembers(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return items
	}
	for _, i := range members {
		items = append(items, i)
	}
	return items
}

func QueryAttributes(b string) map[string]string {
	m, err := nc().HGetAll(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return m
	}
	return m
}
