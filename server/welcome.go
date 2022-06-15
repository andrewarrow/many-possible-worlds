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

	for _, w := range redis.AllWorlds() {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		//buffer = append(buffer, fmt.Sprintf("<a href=\"https://youtube.com/watch?v=%s\">%s</a>", item.Id, item.Title))
		buffer = append(buffer, fmt.Sprintf("<a href=\"/w/%s\">%s</a>", w.Slug, w.Title))

		buffer = append(buffer, "</div>")
		//buffer = append(buffer, "<div>")
		//buffer = append(buffer, item.ViewCount)
		//buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}

	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
