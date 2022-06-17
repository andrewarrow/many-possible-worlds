package server

import (
	"fmt"
	"html/template"
	"many-pw/network"
	"many-pw/parse"
	"many-pw/redis"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var mReferers = map[string]int{}
var mIps = map[string]int{}
var mRoutes = map[string]int{}
var mutex sync.Mutex

func BumpStats(route string, c *gin.Context) {
	ip := c.ClientIP()
	if ip == os.Getenv("MANY_PW_ADMIN_IP") {
		return
	}
	refs := c.Request.Header["Referer"]

	mutex.Lock()
	defer mutex.Unlock()

	if len(refs) > 0 {
		ref := refs[0]
		mReferers[ref]++
	}
	//mIps[ip]++
	//mRoutes[route]++
}

func StatsIndex(c *gin.Context) {
	email, _ := c.Cookie("email")
	if email != os.Getenv("MANY_PW_ADMIN_EMAIL") {
		c.Redirect(http.StatusFound, "/")
		c.Abort()
		return
	}

	cmap := map[string]bool{}
	for _, w := range redis.QueryWorlds() {
		fmt.Println(w)
		fresh := redis.QueryChannelsInSlug(w.Slug, 0)
		for _, f := range fresh {
			if f.ImageUrl == "" {
				cmap[f.Id] = true
			}
		}
		pinned := redis.QueryPinned(w.Slug)
		for _, f := range pinned {
			if f.ImageUrl == "" {
				cmap[f.Id] = true
			}
		}
	}
	json := network.GetChannels(cmap)
	channels := parse.ParseChannelJson(json)
	for _, item := range channels.Items {
		redis.UpdateChannelImage(item.Id, item.Snippet.Thumbnails.Medium.Url)
	}

	body := template.HTML(makeStatsHTML())

	c.HTML(http.StatusOK, "general.tmpl", gin.H{
		"email": "",
		"flash": "",
		"body":  body,
	})
}

func makeStatsHTML() string {
	buffer := []string{}
	mutex.Lock()
	defer mutex.Unlock()
	for k, v := range mReferers {
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("%s<br/>%d", k, v))
		buffer = append(buffer, "</div>")
	}
	/*
		for k, v := range mRoutes {
			buffer = append(buffer, "<div>")
			buffer = append(buffer, fmt.Sprintf("%s<br/>%d", k, v))
			buffer = append(buffer, "</div>")
		}
		for k, v := range mIps {
			buffer = append(buffer, "<div>")
			buffer = append(buffer, fmt.Sprintf("%s<br/>%d", k, v))
			buffer = append(buffer, "</div>")
		}*/
	return strings.Join(buffer, "\n")
}
