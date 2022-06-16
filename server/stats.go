package server

import (
	"fmt"
	"html/template"
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
	mIps[ip]++
	mRoutes[route]++
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
	mutex.Lock()
	defer mutex.Unlock()
	for k, v := range mReferers {
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("%s<br/>%d", k, v))
		buffer = append(buffer, "</div>")
	}
	for k, v := range mRoutes {
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("%s<br/>%d", k, v))
		buffer = append(buffer, "</div>")
	}
	for k, v := range mIps {
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("%s<br/>%d", k, v))
		buffer = append(buffer, "</div>")
	}
	return strings.Join(buffer, "\n")
}
