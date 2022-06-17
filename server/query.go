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
	BumpStats(slug, c)
	body := template.HTML(makeQueryHTML(slug, offsetInt))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"email": "",
		"flash": "",
		"body":  body,
	})
}

func makeQueryHTML(slug string, offset int) string {
	buffer := []string{}
	items := redis.QueryChannelsInSlug(slug, offset)
	pinned := redis.QueryPinned(slug)

	buffer = append(buffer, "<h2>Pinned</h2>")
	for _, gem := range pinned {
		buffer = append(buffer, "<p>")
		buffer = append(buffer, fmt.Sprintf("<a href=\"/c/%s/%s\">%s</a>", slug, gem.Id, gem.Title))

		buffer = append(buffer, "</p>")
	}

	buffer = append(buffer, "<h2>Latest</h2>")
	t := `
<a href="/c/%s/%s">%s</a> with %s sub(s)<br/>%s ago
`
	for _, item := range items {
		buffer = append(buffer, "<p>")

		buffer = append(buffer, fmt.Sprintf(t, slug, item.Id, item.Title, item.SubscriberCount, DeltaAgo(item.PublishedAt)))

		buffer = append(buffer, "</p>")
	}
	buffer = append(buffer, "<p>")
	buffer = append(buffer, fmt.Sprintf("<a href=\"?offset=%d\">Load Next 50</a>", offset+50))
	buffer = append(buffer, "</p>")
	return strings.Join(buffer, "\n")
}
