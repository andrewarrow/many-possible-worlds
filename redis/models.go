package redis

type Video struct {
	Id           string
	Title        string
	PublishedAt  int64
	ChannelTitle string
	ChannelId    string
	ViewCount    string
	Subs         string
}

type Channel struct {
	ViewCount       string
	VideoCount      string
	SubscriberCount string
	Title           string
	Id              string
	PublishedAt     int64
}

type World struct {
	Slug  string
	Title string
	Score int64
}
