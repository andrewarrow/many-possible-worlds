package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideosShow(c *gin.Context) {

	id := c.Param("id")
	video := redis.LoadVideo(id)

	c.HTML(http.StatusOK, "videos_show.tmpl", gin.H{
		"flash": "",
		"email": "",
		"v":     video,
	})
}
