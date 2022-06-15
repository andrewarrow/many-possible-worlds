package server

import (
	"many-pw/redis"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddWorldIndex(c *gin.Context) {
	flash, _ := c.Cookie("flash")
	c.SetCookie("flash", "", 3600, "/", "", false, true)

	c.HTML(http.StatusOK, "add_world.tmpl", gin.H{
		"flash": flash,
		"email": "",
	})
}
func AddWorldSubmit(c *gin.Context) {
	email, _ := c.Cookie("email")
	if email != os.Getenv("MANY_PW_ADMIN_EMAIL") {
		c.SetCookie("flash", "not valid", 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/add-world")
		c.Abort()
		return
	}

	q := strings.TrimSpace(c.PostForm("q"))
	lower := strings.ToLower(q)
	slug := strings.ReplaceAll(lower, " ", "-")
	redis.InsertWorld(q, slug)

	c.Redirect(http.StatusFound, "/")
	c.Abort()

}
