package parse

import (
	"encoding/json"
	"time"
)

type Result struct {
	NextPageToken string   `json:"nextPageToken"`
	Items         []Item   `json:"items"`
	PageInfo      PageInfo `json:"pageInfo"`
}

type PageInfo struct {
	TotalResults int `json:"totalResults"`
}

type Item struct {
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Title       string    `json:"title"`
	PublishedAt time.Time `json:publishedAt"`
}

/*

  "pageInfo": {
    "totalResults": 1000000,
    "resultsPerPage": 25
  },

   "snippet": {
     "publishedAt": "2016-10-12T22:32:12Z",
     "channelId": "UChSpME3QaSFAWK8Hpmg-Dyw",
     "title": "Daily Calm | 10 Minute Mindfulness Meditation | Be Present",

*/

func ParseJson(jsonString string) ([]Item, int) {
	var result Result
	json.Unmarshal([]byte(jsonString), &result)
	return result.Items, result.PageInfo.TotalResults
}
