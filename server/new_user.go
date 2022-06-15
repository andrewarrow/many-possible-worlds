package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserIndex(c *gin.Context) {

	email, _ := c.Cookie("email")
	password, _ := c.Cookie("password")

	c.HTML(http.StatusOK, "new_user.tmpl", gin.H{
		"flash":    "",
		"email":    email,
		"password": password,
	})
}
