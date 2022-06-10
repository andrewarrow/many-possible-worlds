package main

import (
	"fmt"
	"many-pw/network"
	"many-pw/parse"
	"many-pw/redis"
	"time"
)

func main() {
	//words := []string{"meditation", "spirituality", "awakening", "law of attraction", "ego", "self hypnosis", "manifestation", "zen master", "non-duality", "healing", "thoughts"}
	words := []string{"meditation"}
	for _, word := range words {
		QueryYoutubeUpdateRedis(word)
	}
}

type Video struct {
	Id           string
	Title        string
	PublishedAt  int64
	ChannelTitle string
	ChannelId    string
	ViewCount    string
}

type Channel struct {
	ViewCount       string
	VideoCount      string
	SubscriberCount string
}

func QueryYoutubeUpdateRedis(word string) {
	t := time.Now()
	json := network.SearchWord(word)
	if json != "" {

		result := parse.ParseJson(json)

		//fmt.Println(result.NextPageToken, result.PageInfo.TotalResults)
		ids := []string{}
		vmap := map[string]*Video{}
		cmap := map[string]bool{}
		for _, item := range result.Items {
			v := Video{}
			v.Id = item.Id.VideoId
			v.Title = item.Snippet.Title
			v.PublishedAt = item.Snippet.PublishedAt.Unix()
			v.ChannelTitle = item.Snippet.ChannelTitle
			v.ChannelId = item.Snippet.ChannelId
			vmap[v.Id] = &v
			//fmt.Println(item.Id.VideoId, item.Snippet.PublishedAt, item.Snippet.Title)
			ids = append(ids, v.Id)
			cmap[item.Snippet.ChannelId] = true
		}
		json = network.SearchVideos(ids)
		if json != "" {
			stats := parse.ParseStatJson(json)
			for _, stat := range stats.Items {
				v := vmap[stat.Id]
				v.ViewCount = stat.Statistics.ViewCount
				//fmt.Println(stat.Id, stat.Statistics.ViewCount)
			}
		}
		json = network.GetChannels(cmap)
		//ioutil.WriteFile("fname.txt", []byte(json), 0644)
		channels := parse.ParseChannelJson(json)
		channelStats := map[string]Channel{}
		for _, item := range channels.Items {
			c := Channel{}
			c.ViewCount = item.Statistics.ViewCount
			c.VideoCount = item.Statistics.VideoCount
			c.SubscriberCount = item.Statistics.SubscriberCount
			channelStats[item.Id] = c
		}

		for _, v := range vmap {
			c := channelStats[v.ChannelId]
			fmt.Printf("%05s %s\n", v.ViewCount, v.Title)
			fmt.Printf("%05s %s\n", c.SubscriberCount, v.ChannelTitle)
			redis.InsertItem(t.Unix(), v.Id, v.Title, v.ChannelTitle, v.ViewCount, v.ChannelId, c.SubscriberCount)
		}

	}
}
