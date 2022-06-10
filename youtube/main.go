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
	json := network.SearchWord(word)
	if json != "" {

		ioutil.WriteFile("fname.txt", []byte(json), 0644)

		result := parse.ParseJson(json)
		fmt.Println(result.PageInfo.TotalResults)
		ids := []string{}
		for _, item := range result.Items {
			fmt.Println(item.Id.VideoId, item.Snippet.PublishedAt, item.Snippet.Title)
			ids = append(ids, item.Id.VideoId)
		}

		fmt.Println(ids)
		json = network.SearchVideos(ids)
		ioutil.WriteFile("fname2.txt", []byte(json), 0644)
	}
}
