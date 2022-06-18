package parse

import "encoding/json"

type VideoItems struct {
	Items []StatItem `json:"items"`
}

type VideoItem struct {
	Id         string     `json:"id"`
	Statistics Statistics `json:"statistics"`
}

func ParseVideoJson(jsonString string) *VideoItems {
	var result VideoItems
	json.Unmarshal([]byte(jsonString), &result)
	return &result
}
