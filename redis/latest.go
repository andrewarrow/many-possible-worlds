package redis

type Latest struct {
	Image                   string
	ChannelUrl              string
	About                   string
	ExampleVideoId          string
	ExampleVideoPublishedAt int64
	ExampleVideoTitle       string
	SubscriberCount         int64
	ViewCount               int64
	VideoCount              int64
}

func QueryLatest() []Latest {
	items := []Latest{}
	l := Latest{}
	l.Image = "https://yt3.ggpht.com/Nfwla_eYJXBx9Cro4_qaAtadV48BzpVGQ7OTo47yZJF3ExVza2selzvYOyl3SpdMOqM4sGXP=s176-c-k-c0x00ffffff-no-rj"
	l.ChannelUrl = "https://www.youtube.com/channel/UCQWA5jLmOtqymAjc_9vkn3A"
	l.About = "Kelly Hart has a soulful look @ Present Evolution & Ascension. She is looking at reality in a new way. Healing at a time of change, transformation and through Love. New ways of living and creating the world we want."
	l.ExampleVideoId = "HliJl8dXIuI"
	l.ExampleVideoPublishedAt = 1655492657
	l.ExampleVideoTitle = "Health of all the world"
	l.SubscriberCount = 516
	l.ViewCount = 123999
	l.VideoCount = 120
	items = append(items, l)
	return items
}
