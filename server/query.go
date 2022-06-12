package server

import (
	"fmt"
	"html/template"
	"many-pw/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func QueryIndex(c *gin.Context) {

	slug := c.Param("world")
	body := template.HTML(makeQueryHTML(slug))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeQueryHTML(slug string) string {
	buffer := []string{}

	items := redis.QueryDay(slug)

	buffer = append(buffer, "<div class=\"good-links\">")

	for _, item := range items {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("<a href=\"https://youtube.com/watch?v=%s\">%s</a>", item.Id, item.Title))

		buffer = append(buffer, "</div>")
		buffer = append(buffer, "<div>")
		buffer = append(buffer, item.ViewCount)
		buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}

	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
