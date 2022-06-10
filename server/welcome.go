package server

import (
	"fmt"
	"html/template"
	"many-pw/redis"
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

	items := redis.QueryHour()

	for _, item := range items {
		buffer = append(buffer, fmt.Sprintf("<div class=\"item\"><a href=\"https://youtube.com/watch?v=%s\">%s</a></div>", item.Id, item.Id))
	}

	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
