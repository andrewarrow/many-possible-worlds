package redis

type Video struct {
	Id           string
	Title        string
	PublishedAt  int64
	ChannelTitle string
	ChannelId    string
	ViewCount    string
}

type Channel struct {
	ViewCount       string
	VideoCount      string
	SubscriberCount string
}
