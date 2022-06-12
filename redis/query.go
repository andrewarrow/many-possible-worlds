package redis

import (
	"fmt"
	"time"
)

func QueryDay() []Video {
	t := time.Now().In(utc)
	t = t.Add(time.Hour * -1)
	bucket := BucketForDay(t)
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

	vals, err := nc().ZRangeWithScores(ctx, b, 0, -1).Result()

	if err != nil {
		fmt.Println(err)
		return list
	}
	for _, item := range vals {
		fmt.Println(int64(item.Score), item.Member)
		/*
			v := Video{}
			v.Id = item
			m := QueryAttributes(v.Id)
			v.Title = m["title"]
			v.ViewCount = m["view_count"]
			v.ChannelId = m["c_id"]
			list = append(list, v)
		*/
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
