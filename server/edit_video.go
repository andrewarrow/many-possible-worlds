package server

import (
	"many-pw/redis"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func VideoEdit(c *gin.Context) {

	id := c.Param("id")
	video := redis.LoadVideo(id)
	c.HTML(http.StatusOK, "edit_video.tmpl", gin.H{
		"flash": "",
		"v":     video,
		"email": "",
	})
}

func VideoEditSubmit(c *gin.Context) {
	email, _ := c.Cookie("email")
	if email != os.Getenv("MANY_PW_ADMIN_EMAIL") {
		c.SetCookie("flash", "not valid", 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/add-video")
		c.Abort()
		return
	}

	id := c.Param("id")
	highlight := strings.TrimSpace(c.PostForm("h"))
	redis.UpdateVideoHighlight(id, highlight)

	c.Redirect(http.StatusFound, "/videos")
	c.Abort()
}
