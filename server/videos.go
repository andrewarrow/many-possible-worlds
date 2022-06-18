package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideosShow(c *gin.Context) {

	id := c.Param("id")
	video := redis.LoadVideo(id)
	older := redis.FindOlderVideo(video.ChannelId, video.PublishedAt)
	channel := redis.LoadLatest(video.ChannelId)
	redis.UpdateLatest(video.ChannelId, channel.ViewCount)

	c.HTML(http.StatusOK, "videos_show.tmpl", gin.H{
		"flash": "",
		"email": "",
		"v":     video,
		"older": older,
		"c":     channel,
	})
}
