package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}
	command := os.Args[1]
	//argMap := ArgsToMap()

	if command == "v" {
		id := os.Args[2]
		ImportVideo(id)
	} else if command == "c" {
		id := os.Args[2]
		ImportChannel(id)
	} else if command == "transcript" {
		id := os.Args[2]
		ParseTranscript(id)
	}
}
