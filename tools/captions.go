package main

import (
	"io/ioutil"
	"many-pw/network"
)

func ImportCaptions(id string) {
	//json := network.FetchCaptions(id)
	json := network.DownloadCaption(id)
	if json == "" {
		return
	}
	ioutil.WriteFile("fname.txt", []byte(json), 0644)
}
