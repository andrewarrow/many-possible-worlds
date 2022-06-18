package main

import (
	"fmt"
	"many-pw/network"
	"many-pw/parse"
)

func ImportVideo(id string) {
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
