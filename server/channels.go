package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChannelsIndex(c *gin.Context) {

	sort := c.DefaultQuery("sort", "")
	key := "latest"
	if sort == "vc" {
		key = key + "-" + sort
	}
	latest := redis.QueryLatest(key, 50)

	c.HTML(http.StatusOK, "channels.tmpl", gin.H{
		"flash":  "",
		"email":  "",
		"latest": latest,
	})
}
