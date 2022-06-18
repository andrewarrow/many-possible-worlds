package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Latest struct {
	ChannelId               string
	ChannelTitle            string
	ImageUrl                string
	ExampleVideoId          string
	ExampleVideoPublishedAt int64
	ExampleVideoTitle       string
	SubscriberCount         int64
	ViewCount               int64
	VideoCount              int64
}

func LoadLatest(id string) *Latest {
	l := Latest{}
	l.ChannelId = id
	m := QueryAttributes(l.ChannelId)
	l.ImageUrl = m["img"]
	l.ChannelTitle = m["title"]
	l.ExampleVideoId = m["vid"]
	l.ExampleVideoPublishedAt, _ = strconv.ParseInt(m["vat"], 10, 64)
	l.ExampleVideoTitle = m["vt"]
	l.SubscriberCount, _ = strconv.ParseInt(m["subs"], 10, 64)
	l.ViewCount, _ = strconv.ParseInt(m["vc"], 10, 64)
	l.VideoCount, _ = strconv.ParseInt(m["vidc"], 10, 64)
	return &l
}

func QueryLatest(zset string, amount int) []*Latest {
	list := []*Latest{}

	vals, _ := nc().ZRevRangeWithScores(ctx, zset, int64(0), int64(amount-1)).Result()
	for _, item := range vals {
		id := item.Member.(string)
		l := LoadLatest(id)
		//w.Score = int64(item.Score)
		list = append(list, l)
	}

	return list
}

func UpdateLatestVc(l *Latest) {
	zset := "latest-vc"
	rz := redis.Z{Score: float64(l.ViewCount), Member: l.ChannelId}
	nc().ZAdd(ctx, zset, &rz).Err()
}

func UpdateLatest(id string, viewCount int64) {
	if id == "" {
		return
	}
	if viewCount > 99999 {
		return
	}
	zset := "latest"
	rz := redis.Z{Score: float64(time.Now().Unix()), Member: id}
	nc().ZAdd(ctx, zset, &rz).Err()
}

func InsertLatest(l *Latest) {

	zset := "latest"
	rz := redis.Z{Score: float64(time.Now().Unix()), Member: l.ChannelId}
	nc().ZAdd(ctx, zset, &rz).Err()

	zset = "latest-vc"
	rz = redis.Z{Score: float64(l.ViewCount), Member: l.ChannelId}
	nc().ZAdd(ctx, zset, &rz).Err()

	nc().HSet(ctx, l.ChannelId, "title", l.ChannelTitle).Err()
	nc().HSet(ctx, l.ChannelId, "img", l.ImageUrl).Err()
	nc().HSet(ctx, l.ChannelId, "vid", l.ExampleVideoId).Err()
	nc().HSet(ctx, l.ChannelId, "vat", l.ExampleVideoPublishedAt).Err()
	nc().HSet(ctx, l.ChannelId, "vt", l.ExampleVideoTitle).Err()
	nc().HSet(ctx, l.ChannelId, "subs", fmt.Sprintf("%d", l.SubscriberCount)).Err()
	nc().HSet(ctx, l.ChannelId, "vc", fmt.Sprintf("%d", l.ViewCount)).Err()
	nc().HSet(ctx, l.ChannelId, "vidc", fmt.Sprintf("%d", l.VideoCount)).Err()

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	//nc().ExpireAt(ctx, v.Id, expireTime)
	nc().ExpireAt(ctx, l.ChannelId, expireTime)
}

func OldQueryLatest() []Latest {
	items := []Latest{}
	l := Latest{}
	l.ImageUrl = "https://yt3.ggpht.com/Nfwla_eYJXBx9Cro4_qaAtadV48BzpVGQ7OTo47yZJF3ExVza2selzvYOyl3SpdMOqM4sGXP=s176-c-k-c0x00ffffff-no-rj"
	l.ChannelId = "UCQWA5jLmOtqymAjc_9vkn3A"
	//l.About = "Kelly Hart has a soulful look @ Present Evolution & Ascension. She is looking at reality in a new way. Healing at a time of change, transformation and through Love. New ways of living and creating the world we want."
	l.ExampleVideoId = "HliJl8dXIuI"
	l.ExampleVideoPublishedAt = 1655492657
	l.ExampleVideoTitle = "Health of all the world"
	l.SubscriberCount = 516
	l.ViewCount = 123999
	l.VideoCount = 120
	items = append(items, l)

	l.ImageUrl = "https://yt3.ggpht.com/KIbYhYMDIiNMCEXrD9Yr_Gc6HrZUt49ASJ-bsFlwt4lcIfbwBu1DVjzjK-nvjUGMr6mokbaG0iI=s176-c-k-c0x00ffffff-no-rj"
	l.ChannelId = "UCAoZk1bZVpxFJY2pQzh9w8Q"
	//l.About = "Duality & Non-duality - Pointings and ramblings on the unnamable and how it's affecting my relative experience."
	l.ExampleVideoId = "GdB-aI8cKhc"
	l.ExampleVideoPublishedAt = 1655292657
	l.ExampleVideoTitle = "Tuning forks, frequencies and Sangha"
	l.SubscriberCount = 45
	l.ViewCount = 744
	l.VideoCount = 18
	items = append(items, l)

	l.ImageUrl = "https://yt3.ggpht.com/ytc/AKedOLR8DZsyieF9ky6wb-x_mx5nVzR7oPG3OCLa8XG3GA=s176-c-k-c0x00ffffff-no-rj"
	l.ChannelId = "rgold3206"
	//l.About = "Christian / Religious Programming."
	l.ExampleVideoId = "HSI9tlBhu-c"
	l.ExampleVideoPublishedAt = 1655492657
	l.ExampleVideoTitle = "God Chose You"
	l.SubscriberCount = 482
	l.ViewCount = 61554
	l.VideoCount = 180
	items = append(items, l)
	return items
}
