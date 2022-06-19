package server

import (
	"html/template"
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideosShow(c *gin.Context) {

	id := c.Param("id")
	video := redis.LoadVideo(id)
	if video.Title == "" {
		Video404(c)
		return
	}
	prev, next := redis.FindPrevAndNextVideos(id, video.ChannelId)
	channel := redis.LoadChannel(video.ChannelId)
	redis.UpdateLatest(video.ChannelId, channel.ViewCount)

	c.HTML(http.StatusOK, "videos_show.tmpl", gin.H{
		"flash": "",
		"email": "",
		"v":     video,
		"prev":  prev,
		"next":  next,
		"c":     channel,
	})
}

func Video404(c *gin.Context) {
	body := template.HTML("This video not found, but contact us to add it!")

	c.HTML(http.StatusNotFound, "general.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})

}
