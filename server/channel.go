package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChannelShow(c *gin.Context) {

	id := c.Param("id")
	slug := c.Param("slug")

	items := redis.QueryVideosInChannel(id, 0)

	c.HTML(http.StatusOK, "channel.tmpl", gin.H{
		"flash": "",
		"id":    id,
		"items": items,
		"slug":  slug,
	})
}
