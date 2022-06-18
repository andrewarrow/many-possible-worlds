package parse

import "encoding/json"

type VideoItems struct {
	Items []VideoItem `json:"items"`
}

type VideoItem struct {
	Id         string     `json:"id"`
	Statistics Statistics `json:"statistics"`
	Snippet    Snippet    `json:"snippet"`
}

func ParseVideoJson(jsonString string) *VideoItems {
	var result VideoItems
	json.Unmarshal([]byte(jsonString), &result)
	return &result
}
