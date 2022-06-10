package main

import (
	"fmt"
	"youtube/network"
)

func main() {
	word := "meditation"
	fmt.Println("searching for", word)
	json := network.SearchWord(word)
	fmt.Println(json)
}
