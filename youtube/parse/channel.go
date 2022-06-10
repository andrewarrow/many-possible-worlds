package parse

import "encoding/json"

type ChannelItems struct {
	Items []ChannelItem `json:"items"`
}

type ChannelItem struct {
	Id         string            `json:"id"`
	Statistics ChannelStatistics `json:"statistics"`
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
