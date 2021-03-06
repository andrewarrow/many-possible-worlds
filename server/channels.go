package server

import (
	"html/template"
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

func ChannelsShow(c *gin.Context) {

	id := c.Param("id")
	channel := redis.LoadChannel(id)
	if channel.Title == "" {
		Channel404(c)
		return
	}
	redis.UpdateLatest(id, channel.ViewCount)
	c.HTML(http.StatusOK, "channels_show.tmpl", gin.H{
		"flash": "",
		"email": "",
		"c":     channel,
	})
}

func Channel404(c *gin.Context) {
	body := template.HTML("This channel not found, but contact us to add it!")

	c.HTML(http.StatusNotFound, "general.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})

}
