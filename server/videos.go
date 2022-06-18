package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideosShow(c *gin.Context) {

	id := c.Param("id")
	video := redis.LoadVideo(id)
	redis.UpdateLatest(video.ChannelId)
	channel := redis.LoadLatest(video.ChannelId)

	c.HTML(http.StatusOK, "videos_show.tmpl", gin.H{
		"flash": "",
		"email": "",
		"v":     video,
		"c":     channel,
	})
}
