package redis

func AllWorlds() []World {
	w1 := World{}
	w1.Slug = "meditation"
	w1.Title = "meditation"

	w2 := World{}
	w2.Slug = "law-of-attraction"
	w2.Title = "Law of Attraction"

	w3 := World{}
	w3.Slug = "this-just-happened"
	w3.Title = "This JUST Happened"

	items := []World{w1, w2, w3}
	return items
}
