package server

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {

	flash, _ := c.Cookie("flash")
	c.SetCookie("flash", "", 3600, "/", "", false, true)

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"flash": flash,
	})
}

func LoginSubmit(c *gin.Context) {
	/*
		password := c.PostForm("password")
		if password == os.Getenv("MANY_PW_PASSWORD") {
			uuid := pseudoUuid()
			redis.SetAuth(uuid)
			c.SetCookie("auth", uuid, 3600*24*365*10, "/", "", false, true)
		}*/
	cookies := c.PostForm("cookies")
	if cookies != "1" {
		FlashAndReturnLogin(c, "You must agree to cookies")
		return
	}
}

func pseudoUuid() string {

	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
