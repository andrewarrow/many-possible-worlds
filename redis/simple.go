package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Simple() {
	slug := "hi"
	pubzset := fmt.Sprintf("%s-p", slug)
	nc().Del(ctx, pubzset).Err()

	/*
		rz := redis.Z{
			Score:  float64(1200),
			Member: "foo",
		}

		nc().ZAdd(ctx, pubzset, &rz).Err()

		rz = redis.Z{
			Score:  float64(1201),
			Member: "bar",
		}

		nc().ZAdd(ctx, pubzset, &rz).Err()
	*/
	zrb := redis.ZRangeBy{
		Min: "0",
		Max: "1205",
	}

	vals, _ := nc().ZRangeByScore(ctx, pubzset, &zrb).Result()
	fmt.Println(vals)
}
