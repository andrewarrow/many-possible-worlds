package network

import (
	"fmt"
	"os"
	"time"
)

func SearchWord(word string) string {

	t := time.Now().Add(time.Hour * -24)
	pa := t.Format(time.RFC3339)
	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("search?part=snippet&order=date&maxResults=50&q=%s&key=%s&publishedAfter=%s", word, key, pa))
	return json
}
