package main

import (
	"fmt"
	"io/ioutil"
	"youtube/network"
	"youtube/parse"
)

func main() {
	word := "meditation"
	fmt.Println("searching for", word)
	json := network.SearchWord(word, "")
	if json != "" {

		result := parse.ParseJson(json)

		fmt.Println(result.NextPageToken, result.PageInfo.TotalResults)
		ids := []string{}
		cmap := map[string]bool{}
		for _, item := range result.Items {
			fmt.Println(item.Id.VideoId, item.Snippet.PublishedAt, item.Snippet.Title)
			ids = append(ids, item.Id.VideoId)
			cmap[item.Snippet.ChannelId] = true
		}
		json = network.GetChannels(cmap)
		ioutil.WriteFile("fname.txt", []byte(json), 0644)

		fmt.Println(ids)
		json = network.SearchVideos(ids)
		if json != "" {
			stats := parse.ParseStatJson(json)
			for _, stat := range stats.Items {
				fmt.Println(stat.Id, stat.Statistics.ViewCount)
			}
		}
	}
}
