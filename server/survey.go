package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SurveyIndex(c *gin.Context) {

	flash, _ := c.Cookie("flash")
	c.SetCookie("flash", "", 3600, "/", "", false, true)

	c.HTML(http.StatusOK, "survey.tmpl", gin.H{
		"flash": flash,
		"email": "",
	})
}
