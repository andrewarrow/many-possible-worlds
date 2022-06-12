package server

import (
	"fmt"
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

type World struct {
	Slug  string
	Title string
}

func makeWelcomeHTML() string {
	buffer := []string{}

	buffer = append(buffer, "<div class=\"good-links\">")

	w1 := World{}
	w1.Slug = "meditation"
	w1.Title = "meditation"

	w2 := World{}
	w2.Slug = "law-of-attraction"
	w2.Title = "Law of Attraction"

	items := []World{w1, w2}
	for _, w := range items {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		//buffer = append(buffer, fmt.Sprintf("<a href=\"https://youtube.com/watch?v=%s\">%s</a>", item.Id, item.Title))
		buffer = append(buffer, fmt.Sprintf("<a href=\"/pw/%s\">%s</a>", w.Slug, w.Title))

		buffer = append(buffer, "</div>")
		//buffer = append(buffer, "<div>")
		//buffer = append(buffer, item.ViewCount)
		//buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}

	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
