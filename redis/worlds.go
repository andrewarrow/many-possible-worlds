package redis

func AllWorlds() []World {
	w1 := World{}
	w1.Slug = "meditation"
	w1.Title = "meditation"

	w2 := World{}
	w2.Slug = "law-of-attraction"
	w2.Title = "Law of Attraction"

	items := []World{w1, w2}
	return items
}
