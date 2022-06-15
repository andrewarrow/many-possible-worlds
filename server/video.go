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
	redis.UpdateGemCount(slug, id, 1)
	c.Redirect(http.StatusFound, "/w/"+slug)
	c.Abort()
}

func VideoUnGem(c *gin.Context) {
	id := c.Param("id")
	slug := c.Param("slug")
	redis.UpdateGemCount(slug, id, -1)
	c.Redirect(http.StatusFound, "/w/"+slug)
	c.Abort()
}
