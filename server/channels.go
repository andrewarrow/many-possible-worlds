package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChannelsIndex(c *gin.Context) {

	latest := redis.QueryLatest()

	c.HTML(http.StatusOK, "channels.tmpl", gin.H{
		"flash":  "",
		"email":  "",
		"latest": latest,
	})
}
