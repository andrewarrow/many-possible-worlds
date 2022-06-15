package server

import (
	"fmt"
	"many-pw/redis"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {

	flash, _ := c.Cookie("flash")
	c.SetCookie("flash", "", 3600, "/", "", false, true)

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"flash": flash,
		"email": "",
	})
}

func LoginSubmit(c *gin.Context) {
	cookies := c.PostForm("cookies")
	if cookies != "1" {
		FlashAndReturnLogin(c, "You must agree to cookies")
		return
	}
	password := c.PostForm("password")
	email := c.PostForm("email")

	existing := redis.LookupEmail(email)
	if existing != password {
		FlashAndReturnLogin(c, "Not a valid login.")
		return
	}

	c.SetCookie("email", email, 3600*24*365*10, "/", "", false, true)
	c.SetCookie("password", password, 3600*24*365*10, "/", "", false, true)

	c.Redirect(http.StatusFound, "/")
	c.Abort()

}

func pseudoUuid() string {

	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
