package main

import (
	"fmt"
	"html"
	"many-pw/network"
	"many-pw/parse"
	"many-pw/redis"
	"strconv"
)

func ImportChannel(id string) {
	json := network.VideosInChannel(id)
	if json == "" {
		return
	}
	//ioutil.WriteFile("fname.txt", []byte(json), 0644)
	result := parse.ParseJson(json)

	if len(result.Items) < 1 {
		fmt.Println("< 1")
		return
	}

	first := result.Items[0]
	/*
		ImageUrl                string
	*/

	l := redis.Latest{}
	l.ChannelId = first.Snippet.ChannelId
	l.ChannelTitle = html.UnescapeString(first.Snippet.ChannelTitle)
	l.ExampleVideoId = first.Id.VideoId
	l.ExampleVideoPublishedAt = first.Snippet.PublishedAt.Unix()
	l.ExampleVideoTitle = html.UnescapeString(first.Snippet.Title)

	json = network.GetChannel(id)
	if json == "" {
		return
	}
	//ioutil.WriteFile("fname.txt", []byte(json), 0644)
	channels := parse.ParseChannelJson(json)
	if len(channels.Items) < 1 {
		fmt.Println("channels < 1")
		return
	}

	c := channels.Items[0]

	l.ViewCount, _ = strconv.ParseInt(c.Statistics.ViewCount, 10, 64)
	l.VideoCount, _ = strconv.ParseInt(c.Statistics.VideoCount, 10, 64)
	l.SubscriberCount, _ = strconv.ParseInt(c.Statistics.SubscriberCount, 10, 64)
	l.ImageUrl = c.Snippet.Thumbnails.Medium.Url

	redis.InsertLatest(&l)

}
