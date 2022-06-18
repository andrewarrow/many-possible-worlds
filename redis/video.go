package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func LoadVideo(id string) *Video {
	v := Video{}
	m := QueryAttributes(id)
	v.ImageUrl = m["img"]
	v.Title = m["title"]
	v.ChannelId = m["cid"]
	v.PublishedAt, _ = strconv.ParseInt(m["pub"], 10, 64)
	v.Id = id
	return &v
}

func StoreSingleVideo(v *Video) {
	nc().HSet(ctx, v.Id, "title", v.Title).Err()
	nc().HSet(ctx, v.Id, "cid", v.ChannelId).Err()
	nc().HSet(ctx, v.Id, "img", v.ImageUrl).Err()
	nc().HSet(ctx, v.Id, "pub", v.PublishedAt).Err()
	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, v.Id, expireTime)

	AddToVideoSet(v.ChannelId, v.Id, v.PublishedAt)
}

func AddToVideoSet(cid, vid string, pub int64) {
	vidzset := fmt.Sprintf("%s-v", cid)
	rz := redis.Z{Score: float64(pub), Member: vid}
	nc().ZAdd(ctx, vidzset, &rz).Err()
}

func FindPrevAndNextVideos(vid, cid string) (*Video, *Video) {
	vidzset := fmt.Sprintf("%s-v", cid)
	list := []*Video{}
	vals, _ := nc().ZRevRangeWithScores(ctx, vidzset, int64(0), int64(10)).Result()
	for _, item := range vals {
		v := Video{}
		v.Id = item.Member.(string)
		v.PublishedAt = int64(item.Score)
		list = append(list, &v)
	}
	index := 0
	for i, v := range list {
		if v.Id == vid {
			index = i
			break
		}
	}

	if index == 0 {
		return LoadVideo(list[1].Id), nil
	}
	if index == 9 {
		return nil, LoadVideo(list[8].Id)
	}

	return LoadVideo(list[index+1].Id), LoadVideo(list[index-1].Id)
}
