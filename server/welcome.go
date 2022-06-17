package server

import (
	"many-pw/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {

	BumpStats("root", c)

	email, _ := c.Cookie("email")
	password, _ := c.Cookie("password")
	loggedInAs := ""

	existing := redis.LookupEmail(email)
	if existing == password {
		loggedInAs = email
	}
	worlds := redis.QueryWorlds()
	latest := redis.QueryLatest()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash":  "",
		"email":  loggedInAs,
		"worlds": worlds,
		"latest": latest,
	})

}
