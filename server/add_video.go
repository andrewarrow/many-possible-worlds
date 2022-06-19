package server

import (
	"many-pw/api"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddVideo(c *gin.Context) {
	flash, _ := c.Cookie("flash")
	c.SetCookie("flash", "", 3600, "/", "", false, true)

	c.HTML(http.StatusOK, "add_video.tmpl", gin.H{
		"flash": flash,
		"email": "",
	})
}
func AddVideoSubmit(c *gin.Context) {
	email, _ := c.Cookie("email")
	if email != os.Getenv("MANY_PW_ADMIN_EMAIL") {
		c.SetCookie("flash", "not valid", 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/add-video")
		c.Abort()
		return
	}

	id := strings.TrimSpace(c.PostForm("id"))
	api.ImportVideo(id)

	c.Redirect(http.StatusFound, "/videos")
	c.Abort()

}
