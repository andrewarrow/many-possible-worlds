package main

import (
	"math/rand"
	"os"
	"time"

	"many-pw/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server.Serve(os.Args[1])
}
