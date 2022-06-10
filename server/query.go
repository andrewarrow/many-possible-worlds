package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func QueryIndex(c *gin.Context) {

	body := template.HTML(makeQueryHTML())

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeQueryHTML() string {
	buffer := []string{}

	buffer = append(buffer, "<div class=\"good-links\">")

	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
