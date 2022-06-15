package main

import (
	"fmt"
	"many-pw/network"
	"many-pw/parse"
	"many-pw/redis"
)

func main() {
	//words := []string{"meditation", "spirituality", "awakening", "law of attraction", "ego", "self hypnosis", "manifestation", "zen master", "non-duality", "healing", "thoughts"}
	//words := []string{"meditation"}
	for _, w := range redis.QueryWorlds() {
		QueryYoutubeUpdateRedis(w)
	}
}

func QueryYoutubeUpdateRedis(w redis.World) {
	json := network.SearchWord(w.Title)
	if json != "" {

		result := parse.ParseJson(json)

		//fmt.Println(result.NextPageToken, result.PageInfo.TotalResults)
		ids := []string{}
		vmap := map[string]*redis.Video{}
		cmap := map[string]bool{}
		for _, item := range result.Items {
			v := redis.Video{}
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
		/*
			json = network.SearchVideos(ids)
			if json != "" {
				stats := parse.ParseStatJson(json)
				for _, stat := range stats.Items {
					v := vmap[stat.Id]
					v.ViewCount = stat.Statistics.ViewCount
					//fmt.Println(stat.Id, stat.Statistics.ViewCount)
				}
			}*/
		json = network.GetChannels(cmap)
		//ioutil.WriteFile("fname.txt", []byte(json), 0644)
		channels := parse.ParseChannelJson(json)
		channelStats := map[string]redis.Channel{}
		for _, item := range channels.Items {
			c := redis.Channel{}
			c.ViewCount = item.Statistics.ViewCount
			c.VideoCount = item.Statistics.VideoCount
			c.SubscriberCount = item.Statistics.SubscriberCount
			channelStats[item.Id] = c
		}

		for _, v := range vmap {
			c := channelStats[v.ChannelId]
			fmt.Printf("%05s %s\n", v.ViewCount, v.Title)
			fmt.Printf("%05s %s\n", c.SubscriberCount, v.ChannelTitle)
			redis.InsertItem(v, c.SubscriberCount, w.Slug)
		}

	}
}
