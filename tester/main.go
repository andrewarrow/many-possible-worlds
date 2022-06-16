package main

import (
	"fmt"
	"io/ioutil"
	"many-pw/network"
)

func main() {
	cid := "UCHEoWlkGf-nSVdNk9b6SR0g"
	cmap := map[string]bool{}
	cmap[cid] = true
	json := network.GetChannels(cmap)
	ioutil.WriteFile("fname.txt", []byte(json), 0644)
	fmt.Println(cid)
}
