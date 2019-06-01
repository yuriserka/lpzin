package routers

import (
	"fmt"

	"github.com/yuriserka/lpzin/api/controllers"
	"github.com/yuriserka/lpzin/api/utils"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	conn, err := utils.ConnDB()
	if err != nil {
		panic(fmt.Sprintf("db: %v", err))
	}
	defer conn.Close()

	homeRoutes()
	userRoutes()
}

func Run() {
	router.Run(":8080")
}

func homeRoutes() {
	router.GET("/", controllers.GetHomePage)
}

func userRoutes() {
	router.POST("/usuarios", controllers.PostUsers)
	router.GET("/usuarios", controllers.GetAllUsers)
	router.GET("/usuarios/:userID", controllers.GetUser)
	router.POST("/chat", controllers.PostChat)
	router.GET("/chat/:chatID", controllers.GetChat)
	router.POST("/chat/:chatID", controllers.PostInChat)
}
