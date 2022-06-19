package main

import (
	"many-pw/api"
	"many-pw/redis"
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
		api.ImportVideo(id)
	} else if command == "transcript" {
		id := os.Args[2]
		ParseTranscript(id)
	} else if command == "loop" {
		latest := redis.QueryLatest("latest", 50)
		for _, l := range latest {
			api.ImportVideo(l.ExampleVideoId)
		}
	}
}
