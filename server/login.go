package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"flash": "",
	})
}
