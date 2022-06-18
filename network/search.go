package network

import (
	"fmt"
	"net/url"
	"os"
	"time"
)

func SearchWord(q string) string {

	t := time.Now().Add(time.Hour * -1)
	pa := t.Format(time.RFC3339)
	key := os.Getenv("YOUTUBE_KEY")
	url := fmt.Sprintf("search?part=snippet&order=date&maxResults=50&q=%s&key=%s&publishedAfter=%s", url.QueryEscape(q), key, pa)
	json := DoGet(url)
	return json
}
