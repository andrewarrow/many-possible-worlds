package network

import (
	"fmt"
	"os"
	"strings"
)

func FetchVideos(ids []string) string {

	key := os.Getenv("YOUTUBE_KEY")
	list := strings.Join(ids, ",")
	json := DoGet(fmt.Sprintf("videos?key=%s&fields=items(id(*),snippet(*),statistics(viewCount))&part=id,snippet,statistics&id=%s", key, list))
	return json
}

func FetchVideo(id string) string {
	items := []string{id}
	return FetchVideos(items)
}

func FetchCaptions(id string) string {

	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("captions?key=%s&part=id,snippet&videoId=%s", key, id))
	return json
}

func DownloadCaption(id string) string {

	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("captions/%s?key=%s", id, key))
	return json
}
