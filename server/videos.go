package server

import (
	"fmt"
	"html/template"
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideosIndex(c *gin.Context) {

	key := "latest-vid"
	latest := redis.QueryLatestVideos(key, 50)

	c.HTML(http.StatusOK, "videos.tmpl", gin.H{
		"flash":  "",
		"email":  "",
		"latest": latest,
	})
}

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

	og := NewOpenGraph(channel.Title)
	og.Title = video.Title
	og.Url = fmt.Sprintf("video/%s", video.Id)
	og.Description = "Wake up to Non-Duality and all the possible worlds."
	og.Image = video.ImageUrl

	c.HTML(http.StatusOK, "videos_show.tmpl", gin.H{
		"flash": "",
		"email": "",
		"og":    og,
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
