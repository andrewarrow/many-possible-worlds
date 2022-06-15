package server

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	router := gin.Default()

	prefix := ""
	router.Static("/assets", prefix+"assets")
	router.GET("/", WelcomeIndex)
	router.GET("/login", LoginIndex)
	router.GET("/how-it-works", HowIndex)
	router.GET("/welcome-new-user", NewUserIndex)
	router.GET("/add-world", AddWorldIndex)
	router.GET("/stats", StatsIndex)
	router.GET("/w/:world", QueryIndex)
	router.GET("/c/:slug/:id", ChannelShow)
	router.GET("/v/:slug/:id", VideoShow)
	router.POST("/c/:slug/:id/gem", ChannelGem)
	router.POST("/c/:slug/:id/ungem", ChannelUnGem)
	router.POST("/v/:slug/:id/gem", VideoGem)
	router.POST("/v/:slug/:id/ungem", VideoUnGem)
	router.POST("/login", LoginSubmit)
	router.POST("/logout", LogoutSubmit)
	router.POST("/register", RegisterSubmit)
	router.POST("/add-world", AddWorldSubmit)
	router.NoRoute(NotFoundIndex)

	AddTemplates(router, prefix)
	go router.Run(fmt.Sprintf(":%s", port))

	for {
		time.Sleep(time.Second)
	}

}

func AddTemplates(r *gin.Engine, prefix string) {
	fm := template.FuncMap{
		"mod":    func(i, j int) bool { return i%j == 0 },
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob(prefix + "templates/*.tmpl")
}
