package redis

type Video struct {
	Id           string
	Title        string
	PublishedAt  int64
	ChannelTitle string
	ChannelId    string
	ViewCount    string
	Subs         string
	ImageUrl     string
}

type Channel struct {
	ViewCount               int64
	VideoCount              int64
	SubscriberCount         int64
	Title                   string
	Id                      string
	PublishedAt             int64
	ImageUrl                string
	ExampleVideoId          string
	ExampleVideoPublishedAt int64
	ExampleVideoTitle       string
}

type World struct {
	Slug  string
	Title string
	Score int64
}
