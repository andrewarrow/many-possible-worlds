package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfinityIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "infinity.tmpl", gin.H{
		"flash": "",
		"email": "",
	})
}
