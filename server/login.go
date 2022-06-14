package server

import (
	"fmt"
	"many-pw/redis"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"flash": "",
	})
}

func LoginSubmit(c *gin.Context) {
	//email := c.FormPost("email")
	password := c.Param("password")
	if password == os.Getenv("MANY_PW_PASSWORD") {
		uuid := pseudoUuid()
		redis.SetAuth(uuid)
		c.SetCookie("auth", uuid, 3600*24*365*10, "/", "", false, true)
	}
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

func pseudoUuid() string {

	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
