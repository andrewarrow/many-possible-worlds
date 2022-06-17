package parse

import (
	"encoding/json"
	"time"
)

type ChannelItems struct {
	Items []ChannelItem `json:"items"`
}

type ChannelItem struct {
	Id         string            `json:"id"`
	Statistics ChannelStatistics `json:"statistics"`
	Snippet    ChannelSnippet    `json:""snippet`
}

type ChannelSnippet struct {
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	DefaultLanguage string     `json:"defaultLanguage"`
	PublishedAt     time.Time  `json:publishedAt"`
	Thumbnails      Thumbnails `json:"thumbnails"`
	CustomUrl       string     `json:"customUrl"`
	Country         string     `json:"country"`
}

type ChannelStatistics struct {
	ViewCount       string `json:"viewCount"`
	VideoCount      string `json:"videoCount"`
	SubscriberCount string `json:"subscriberCount"`
}

func ParseChannelJson(jsonString string) *ChannelItems {
	var result ChannelItems
	json.Unmarshal([]byte(jsonString), &result)
	return &result
}
