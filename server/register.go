package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterSubmit(c *gin.Context) {
	cookies := c.PostForm("cookies")
	if cookies != "1" {
		FlashAndReturnLogin(c, "You must agree to cookies")
		return
	}
	email := strings.TrimSpace(c.PostForm("email"))
	if len(email) < 7 || strings.Index(email, "@") == -1 || strings.Index(email, ".") == -1 {
		FlashAndReturnLogin(c, "Email is not valid.")
		return
	}

}

func FlashAndReturnLogin(c *gin.Context, flash string) {
	c.SetCookie("flash", flash, 3600, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}
