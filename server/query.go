package server

import (
	"fmt"
	"html/template"
	"many-pw/redis"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func QueryIndex(c *gin.Context) {

	offset := c.DefaultQuery("offset", "0")
	offsetInt, _ := strconv.Atoi(offset)
	slug := c.Param("world")
	BumpStats(slug, c.ClientIP())
	body := template.HTML(makeQueryHTML(slug, offsetInt))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"email": "",
		"flash": "",
		"body":  body,
	})
}

func makeQueryHTML(slug string, offset int) string {
	buffer := []string{}
	items := redis.QueryChannelsInSlug(slug, 0)
	pinned := redis.QueryPinned(slug)
	buffer = append(buffer, "<div class=\"good-links\">")

	buffer = append(buffer, "<h2>Pinned</h2>")
	for _, gem := range pinned {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("<a href=\"/c/%s/%s\">%s</a>", slug, gem.Id, gem.Title))

		buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}

	buffer = append(buffer, "<h2>Latest</h2>")
	for _, item := range items {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("<a href=\"/c/%s/%s\">%s</a>", slug, item.Id, item.Title))

		buffer = append(buffer, "</div>")
		buffer = append(buffer, "<div class=\"small\">")
		buffer = append(buffer, fmt.Sprintf("%s sub(s)",
			item.SubscriberCount))
		buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}
	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
