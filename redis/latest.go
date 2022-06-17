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

	l.Image = "https://yt3.ggpht.com/KIbYhYMDIiNMCEXrD9Yr_Gc6HrZUt49ASJ-bsFlwt4lcIfbwBu1DVjzjK-nvjUGMr6mokbaG0iI=s176-c-k-c0x00ffffff-no-rj"
	l.ChannelUrl = "https://www.youtube.com/channel/UCAoZk1bZVpxFJY2pQzh9w8Q"
	l.About = "Duality & Non-duality - Pointings and ramblings on the unnamable and how it's affecting my relative experience."
	l.ExampleVideoId = "GdB-aI8cKhc"
	l.ExampleVideoPublishedAt = 1655292657
	l.ExampleVideoTitle = "Tuning forks, frequencies and Sangha"
	l.SubscriberCount = 45
	l.ViewCount = 744
	l.VideoCount = 18
	items = append(items, l)

	l.Image = "https://yt3.ggpht.com/ytc/AKedOLR8DZsyieF9ky6wb-x_mx5nVzR7oPG3OCLa8XG3GA=s176-c-k-c0x00ffffff-no-rj"
	l.ChannelUrl = "https://www.youtube.com/user/rgold3206"
	l.About = "Christian / Religious Programming."
	l.ExampleVideoId = "HSI9tlBhu-c"
	l.ExampleVideoPublishedAt = 1655492657
	l.ExampleVideoTitle = "God Chose You"
	l.SubscriberCount = 482
	l.ViewCount = 61554
	l.VideoCount = 180
	items = append(items, l)
	return items
}
