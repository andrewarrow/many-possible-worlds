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

	cookieAuth, _ := c.Cookie("auth")
	redisAuth := redis.GetAuth()
	auth := ""
	if redisAuth != "" && cookieAuth == redisAuth {
		auth = "auth"
	}

	c.HTML(http.StatusOK, "channel.tmpl", gin.H{
		"flash": "",
		"email": "",
		"id":    id,
		"items": items,
		"auth":  auth,
		"slug":  slug,
	})
}

func ChannelGem(c *gin.Context) {
	id := c.Param("id")
	slug := c.Param("slug")
	redis.UpdateGemCount(slug, id, 1)
	c.Redirect(http.StatusFound, "/w/"+slug)
	c.Abort()
}
