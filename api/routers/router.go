package routers

import (
	"github.com/yuriserka/lpzin/api/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	router.GET("/", controllers.GetHomePage)

	router.POST("/usuarios", controllers.PostUsers)
	router.GET("/usuarios", controllers.GetUsers)
	router.GET("/usuarios/:userID", controllers.ViewUser)
}

func Run() {
	router.Run(":8080")
}
