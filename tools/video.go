package main

import (
	"fmt"
	"html"
	"many-pw/network"
	"many-pw/parse"
	"many-pw/redis"
)

func ImportChannelViaVideoId(id string) {
	json := network.FetchVideo(id)
	if json == "" {
		return
	}
	//ioutil.WriteFile("fname.txt", []byte(json), 0644)
	vs := parse.ParseVideoJson(json)
	v := vs.Items[0]

	fmt.Println(v.Snippet.Title)
	fmt.Println(v.Snippet.ChannelTitle)
	fmt.Println(v.Snippet.ChannelId)
	ImportChannel(v.Snippet.ChannelId)
}

func ImportSingleVideo(id string) {
	json := network.FetchVideo(id)
	if json == "" {
		return
	}
	//ioutil.WriteFile("fname.txt", []byte(json), 0644)
	vs := parse.ParseVideoJson(json)
	item := vs.Items[0]

	fmt.Println(item.Snippet.Title)

	v := redis.Video{}
	v.Id = id
	v.Title = html.UnescapeString(item.Snippet.Title)
	v.PublishedAt = item.Snippet.PublishedAt.Unix()
	v.ChannelTitle = html.UnescapeString(item.Snippet.ChannelTitle)
	v.ChannelId = item.Snippet.ChannelId
	v.ImageUrl = item.Snippet.Thumbnails.Medium.Url

	redis.StoreSingleVideo(&v)
}
