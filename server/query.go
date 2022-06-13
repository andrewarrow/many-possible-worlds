package server

import (
	"fmt"
	"html/template"
	"many-pw/redis"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func QueryIndex(c *gin.Context) {

	offset := c.DefaultQuery("offset", "0")
	offsetInt, _ := strconv.Atoi(offset)
	slug := c.Param("world")
	body := template.HTML(makeQueryHTML(slug, offsetInt))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeQueryHTML(slug string, offset int) string {
	buffer := []string{}

	count := redis.QueryDayCount(slug)
	items, cmap := redis.QueryDay(slug)
	gitems, gcmap := redis.QueryDayGems(slug)

	buffer = append(buffer, "<div class=\"good-links\">")

	buffer = append(buffer, "<h2>Gems</h2>")
	for _, item := range gitems {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("<a href=\"/v/%s/%s\">%s</a>", slug, item.Id, item.Title))

		buffer = append(buffer, "</div>")
		buffer = append(buffer, "<div class=\"small\">")
		buffer = append(buffer, fmt.Sprintf("by %s with %s sub(s)", gcmap[item.ChannelId].Title, gcmap[item.ChannelId].SubscriberCount))
		buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}
	buffer = append(buffer, "<h2>Fresh</h2>")
	buffer = append(buffer, fmt.Sprintf("<h5>%d</h5>", count))
	for _, item := range items {
		buffer = append(buffer, "<div class=\"item\">")
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("<a href=\"/v/%s/%s\">%s</a>", slug, item.Id, item.Title))

		buffer = append(buffer, "</div>")
		buffer = append(buffer, "<div class=\"small\">")
		buffer = append(buffer, item.ViewCount)
		t := time.Unix(item.PublishedAt, 0)
		buffer = append(buffer, fmt.Sprintf("view(s) %s</div>", t.Format(time.RFC850)))
		buffer = append(buffer, "<div class=\"small\">")
		buffer = append(buffer, fmt.Sprintf("by %s with %s sub(s)", cmap[item.ChannelId].Title, cmap[item.ChannelId].SubscriberCount))
		buffer = append(buffer, "</div>")
		buffer = append(buffer, "</div>")
	}

	buffer = append(buffer, "</div>")
	buffer = append(buffer, "<div>")
	buffer = append(buffer, fmt.Sprintf("<a href=\"?offset=%d\">Load Next 50</a>", 50))
	buffer = append(buffer, "</div>")
	return strings.Join(buffer, "\n")
}
