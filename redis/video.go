package redis

import "time"

func LoadVideo(id string) *Video {
	v := Video{}
	m := QueryAttributes(id)
	v.ImageUrl = m["img"]
	v.Title = m["title"]
	v.ChannelId = m["cid"]
	v.Id = id
	return &v
}

func StoreSingleVideo(v *Video) {
	nc().HSet(ctx, v.Id, "title", v.Title).Err()
	nc().HSet(ctx, v.Id, "cid", v.ChannelId).Err()
	nc().HSet(ctx, v.Id, "img", v.ImageUrl).Err()
	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, v.Id, expireTime)
}
