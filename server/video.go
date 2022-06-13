package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoShow(c *gin.Context) {

	id := c.Param("id")

	c.HTML(http.StatusOK, "video.tmpl", gin.H{
		"flash": "",
		"id":    id,
	})
}
