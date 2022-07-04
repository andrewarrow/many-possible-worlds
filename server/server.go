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

	router.Static("/assets", "assets")
	router.GET("/", WelcomeIndex)
	router.GET("/infinity", InfinityIndex)
	router.GET("/login", LoginIndex)
	router.GET("/survey", SurveyIndex)
	router.GET("/channels", ChannelsIndex)
	router.GET("/channel/:id", ChannelsShow)
	router.GET("/videos", VideosIndex)
	router.GET("/video/:id", VideosShow)
	router.GET("/video/:id/edit", VideoEdit)
	router.GET("/how-it-works", HowIndex)
	router.GET("/welcome-new-user", NewUserIndex)
	router.GET("/add-video", AddVideo)
	router.GET("/add-world", AddWorld)
	router.GET("/stats", StatsIndex)
	router.GET("/w/:world", WorldIndex)
	router.GET("/c/:slug/", ChannelIndex)
	router.GET("/c/:slug/:id", ChannelSlugShow)
	router.GET("/v/:slug/:id", VideoSlugShow)
	router.POST("/c/:slug/:id/gem", ChannelGem)
	router.POST("/c/:slug/:id/ungem", ChannelUnGem)
	router.POST("/v/:slug/:id/gem", VideoGem)
	router.POST("/v/:slug/:id/ungem", VideoUnGem)
	router.POST("/login", LoginSubmit)
	router.POST("/logout", LogoutSubmit)
	router.POST("/register", RegisterSubmit)
	router.POST("/add-world", AddWorldSubmit)
	router.POST("/add-video", AddVideoSubmit)
	router.POST("/video/:id/edit", VideoEditSubmit)
	// TODO add way to submit specific videos like
	// https://www.youtube.com/watch?v=kFBFFOGwJ5w
	// The Key to Graduating 3rd Density (Law of One) ft. Kyle Cease
	// Both ACIM and the Law of one say that the 1-way ticket to salvation from this realm is forgiveness of everything. But what does forgiveness really mean? This is how most people see forgiveness wrongly, as an act of victimization instead of the ultimate act of self-empowerment and self-love that it is.
	router.NoRoute(NotFoundIndex)

	AddTemplates(router)
	go router.Run(fmt.Sprintf(":%s", port))

	for {
		time.Sleep(time.Second)
	}

}

func AddTemplates(r *gin.Engine) {
	fm := template.FuncMap{
		"mod":      func(i, j int) bool { return i%j == 0 },
		"tokens":   func(s string, i int) string { return strings.Split(s, ".")[i] },
		"deltaAgo": func(i int64) string { return DeltaAgo(i) },
		"add":      func(i, j int) int { return i + j },
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob("templates/*.tmpl")
}
