package server

import (
	"many-pw/redis"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WorldIndex(c *gin.Context) {

	offset := c.DefaultQuery("offset", "0")
	offsetInt, _ := strconv.Atoi(offset)
	slug := c.Param("world")
	BumpStats(slug, c)

	fresh := redis.QueryChannelsInSlug(slug, offsetInt)
	pinned := redis.QueryPinned(slug)

	c.HTML(http.StatusOK, "world.tmpl", gin.H{
		"email":  "",
		"flash":  "",
		"offset": offsetInt + 50,
		"pinned": pinned,
		"fresh":  fresh,
		"world":  slug,
	})
}
