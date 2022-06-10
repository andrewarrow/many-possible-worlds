package main

import (
	"fmt"
	"youtube/network"
	"youtube/parse"
)

func main() {
	word := "meditation"
	fmt.Println("searching for", word)
	json := network.SearchWord(word)
	if json != "" {
		list, total := parse.ParseJson(json)
		fmt.Println(total)
		for _, item := range list {
			fmt.Println(item.Snippet.PublishedAt, item.Snippet.Title)
		}
	}
}
