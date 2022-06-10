package network

import (
	"fmt"
	"os"
	"strings"
)

func GetChannel(id string) string {

	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("channels?part=snippet,contentDetails,statistics&id=%s&key=%s", id, key))
	return json
}

func GetChannels(cmap map[string]bool) string {

	key := os.Getenv("YOUTUBE_KEY")

	ids := []string{}
	for k, _ := range cmap {
		ids = append(ids, k)
	}

	list := strings.Join(ids, ",")
	json := DoGet(fmt.Sprintf("channels?part=snippet,contentDetails,statistics&id=%s&key=%s", list, key))
	return json
}
