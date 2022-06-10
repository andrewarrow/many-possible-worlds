package network

import (
	"fmt"
	"os"
)

func SearchWord(word string) string {

	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("search?part=snippet&maxResults=25&q=%s&key=%s", word, key))
	return json
}
