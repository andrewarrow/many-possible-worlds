package server

import (
	"many-pw/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ChannelShow(c *gin.Context) {

	id := c.Param("id")
	slug := c.Param("slug")
	BumpStats(slug+"/"+id, c)

	items := redis.QueryVideosInChannel(id, 0)

	auth := ""
	if ModInWorld(c) {
		auth = "auth"
	}

	c.HTML(http.StatusOK, "channel.tmpl", gin.H{
		"flash": "",
		"email": "",
		"id":    id,
		"items": items,
		"auth":  auth,
		"slug":  slug,
	})
}

func ModInWorld(c *gin.Context) bool {
	slug := c.Param("slug")
	email, _ := c.Cookie("email")
	password, _ := c.Cookie("password")

	redisAuth := redis.GetAuth(email, password)
	ok := false
	if redisAuth == "all" {
		ok = true
	} else {
		tokens := strings.Split(redisAuth, ",")
		m := map[string]bool{}
		for _, t := range tokens {
			m[t] = true
		}
		if m[slug] {
			ok = true
		}
	}
	return ok
}

func ChannelGem(c *gin.Context) {
	slug := c.Param("slug")
	if !ModInWorld(c) {
		c.Redirect(http.StatusFound, "/w/"+slug)
		c.Abort()
		return
	}
	id := c.Param("id")
	redis.UpdateGemCount(slug, id, 1)
	c.Redirect(http.StatusFound, "/w/"+slug)
	c.Abort()
}
func ChannelUnGem(c *gin.Context) {
	slug := c.Param("slug")
	if !ModInWorld(c) {
		c.Redirect(http.StatusFound, "/w/"+slug)
		c.Abort()
		return
	}
	id := c.Param("id")
	redis.UpdateGemCount(slug, id, -1)
	c.Redirect(http.StatusFound, "/w/"+slug)
	c.Abort()
}
