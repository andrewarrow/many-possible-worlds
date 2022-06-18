package main

import (
	"fmt"
	"io/ioutil"
	"many-pw/network"
	"many-pw/parse"
)

func ImportVideo(id string) {
	json := network.FetchVideo(id)
	if json == "" {
		return
	}
	ioutil.WriteFile("fname.txt", []byte(json), 0644)
	v := parse.ParseVideoJson(json)
	fmt.Println(v)

}
