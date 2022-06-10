package network

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func SearchWord(word string) string {

	t := time.Now().Add(time.Hour * -1)
	pa := t.Format(time.RFC3339)
	key := os.Getenv("YOUTUBE_KEY")
	url := fmt.Sprintf("search?part=snippet&order=date&maxResults=50&q=%s&key=%s&publishedAfter=%s", word, key, pa)
	json := DoGet(url)
	return json
}

func SearchVideos(ids []string) string {

	key := os.Getenv("YOUTUBE_KEY")
	list := strings.Join(ids, ",")
	json := DoGet(fmt.Sprintf("videos?key=%s&fields=items(id(*),snippet(tags),statistics(viewCount))&part=id,snippet,statistics&id=%s", key, list))
	return json
}
