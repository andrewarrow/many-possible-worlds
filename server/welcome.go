package server

import (
	"fmt"
	"html/template"
	"many-pw/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {

	BumpStats("root", c)
	body := template.HTML(makeWelcomeHTML())

	email, _ := c.Cookie("email")
	password, _ := c.Cookie("password")
	loggedInAs := ""

	existing := redis.LookupEmail(email)
	if existing == password {
		loggedInAs = email
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"email": loggedInAs,
		"body":  body,
	})
}

func makeWelcomeHTML() string {
	buffer := []string{}

	for _, w := range redis.QueryWorlds() {
		buffer = append(buffer, "<p>")
		//buffer = append(buffer, fmt.Sprintf("<a href=\"https://youtube.com/watch?v=%s\">%s</a>", item.Id, item.Title))
		buffer = append(buffer, fmt.Sprintf("<a href=\"/w/%s\">%s</a>", w.Slug, w.Title))

		buffer = append(buffer, "</p>")
	}

	return strings.Join(buffer, "\n")
}
