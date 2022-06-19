package redis

import "strconv"

func LoadChannel(id string) *Channel {
	c := Channel{}
	c.Id = id
	m := QueryAttributes(id)
	c.ImageUrl = m["img"]
	c.Title = m["title"]
	c.ExampleVideoId = m["vid"]
	c.ExampleVideoPublishedAt, _ = strconv.ParseInt(m["vat"], 10, 64)
	c.ExampleVideoTitle = m["vt"]
	c.SubscriberCount, _ = strconv.ParseInt(m["subs"], 10, 64)
	c.ViewCount, _ = strconv.ParseInt(m["vc"], 10, 64)
	c.VideoCount, _ = strconv.ParseInt(m["vidc"], 10, 64)
	return &c
}
