package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChannelShow(c *gin.Context) {

	id := c.Param("id")
	slug := c.Param("slug")

	c.HTML(http.StatusOK, "channel.tmpl", gin.H{
		"flash": "",
		"id":    id,
		"slug":  slug,
	})
}
