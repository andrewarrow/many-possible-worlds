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

	w4 := World{}
	w4.Slug = "disaster"
	w4.Title = "disaster"

	w5 := World{}
	w5.Slug = "collapse"
	w5.Title = "collapse"

	items := []World{w1, w2, w3, w4, w5}
	//items := []World{w1}
	return items
}
