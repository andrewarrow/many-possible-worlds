package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HowIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "how.tmpl", gin.H{
		"flash": "",
		"email": "",
	})
}
