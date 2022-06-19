package main

import (
	"fmt"
	"html"
	"many-pw/network"
	"many-pw/parse"
	"many-pw/redis"
	"strconv"
)

func ImportChannel(v *redis.Video) {

	c := redis.Channel{}
	c.Id = v.ChannelId
	c.Title = html.UnescapeString(v.ChannelTitle)
	c.ExampleVideoId = v.Id
	c.ExampleVideoPublishedAt = v.PublishedAt
	c.ExampleVideoTitle = html.UnescapeString(v.Title)

	json := network.GetChannel(c.Id)
	if json == "" {
		return
	}
	//ioutil.WriteFile("fname.txt", []byte(json), 0644)
	channels := parse.ParseChannelJson(json)
	if len(channels.Items) < 1 {
		fmt.Println("channels < 1")
		return
	}

	cjson := channels.Items[0]

	c.ViewCount, _ = strconv.ParseInt(cjson.Statistics.ViewCount, 10, 64)
	c.VideoCount, _ = strconv.ParseInt(cjson.Statistics.VideoCount, 10, 64)
	c.SubscriberCount, _ = strconv.ParseInt(cjson.Statistics.SubscriberCount, 10, 64)
	c.ImageUrl = cjson.Snippet.Thumbnails.Medium.Url

	redis.InsertLatest(&c)
}
