package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HowIndex(c *gin.Context) {

	BumpStats("how", c)
	c.HTML(http.StatusOK, "how.tmpl", gin.H{
		"flash": "",
		"email": "",
	})
}
