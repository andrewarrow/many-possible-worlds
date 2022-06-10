package parse

import "encoding/json"

type Result struct {
	NextPageToken string `json:"nextPageToken"`
	Items         []Item `json:"items"`
}

type Item struct {
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Title string `json:"title"`
}

/*
   "snippet": {
     "publishedAt": "2016-10-12T22:32:12Z",
     "channelId": "UChSpME3QaSFAWK8Hpmg-Dyw",
     "title": "Daily Calm | 10 Minute Mindfulness Meditation | Be Present",

*/

func ParseJson(jsonString string) []Item {
	var result Result
	json.Unmarshal([]byte(jsonString), &result)
	return result.Items
}
