package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoShow(c *gin.Context) {

	id := c.Param("id")
	slug := c.Param("slug")

	c.HTML(http.StatusOK, "video.tmpl", gin.H{
		"flash": "",
		"id":    id,
		"slug":  slug,
	})
}

func VideoGem(c *gin.Context) {
	id := c.Param("id")
	slug := c.Param("slug")
	redis.UpdateGemCount(slug, id)
	c.Redirect(http.StatusFound, "/pw/"+slug)
	c.Abort()
}
