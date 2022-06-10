package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {

	body := template.HTML(makeWelcomeHTML())

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeWelcomeHTML() string {
	buffer := []string{}

	buffer = append(buffer, "<div class=\"good-links\">")

	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
