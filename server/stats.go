package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var stats = []string{}

func BumpStats(route string, c *gin.Context) {
	ip := c.ClientIP()
	if ip == os.Getenv("MANY_PW_ADMIN_IP") {
		return
	}
	refs := c.Request.Header["Referer"]
	ref := ""
	if len(refs) > 0 {
		ref = refs[0]
	}
	if len(stats) > 1000 {
		stats = []string{}
	}
	payload := fmt.Sprintf("%d/%s/%s/%s", time.Now().Unix(), ip, route, ref)
	stats = append([]string{payload}, stats...)
}

func StatsIndex(c *gin.Context) {
	email, _ := c.Cookie("email")
	if email != os.Getenv("MANY_PW_ADMIN_EMAIL") {
		c.Redirect(http.StatusFound, "/")
		c.Abort()
		return
	}

	body := template.HTML(makeStatsHTML())

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"email": "",
		"flash": "",
		"body":  body,
	})
}

func makeStatsHTML() string {
	buffer := []string{}
	for _, item := range stats {
		buffer = append(buffer, "<div>")
		buffer = append(buffer, item)
		buffer = append(buffer, "</div>")
	}
	return strings.Join(buffer, "\n")
}
