package parse

import "encoding/json"

type StatItems struct {
	Items []StatItem `json:"items"`
}

type StatItem struct {
	Id         string     `json:"id"`
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	ViewCount string `json:"viewCount"`
}

func ParseStatJson(jsonString string) *StatItems {
	var result StatItems
	json.Unmarshal([]byte(jsonString), &result)
	return &result
}

/*

{
  "items": [
    {
      "id": "qQxcQqOKHSA",
      "snippet": {
        "tags": [
          "Mother Meera"
        ]
      },
      "statistics": {
        "viewCount": "1141"
      }
    },
*/
