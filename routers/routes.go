package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/controllers"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.LoadHTMLGlob("views/*")
	// Handle the index route
	router.GET("/", controllers.ShowIndexPage)
	router.GET("/article/view/:article_id", controllers.GetArticle)
	router.Static("/static/", "./static")
}

func Run() {
	router.Run()
}
