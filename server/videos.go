package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideosShow(c *gin.Context) {

	id := c.Param("id")
	video := redis.LoadVideo(id)
	prev, next := redis.FindPrevAndNextVideos(id, video.ChannelId)
	channel := redis.LoadLatest(video.ChannelId)
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
