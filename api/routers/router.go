package routers

import (
	"fmt"
	"log"

	"github.com/yuriserka/lpzin/api/common"
	"github.com/yuriserka/lpzin/api/controllers"
	"github.com/yuriserka/lpzin/schema"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	db, err := common.ConnDB()
	if err != nil {
		log.Panic(fmt.Sprintf("db: %v", err))
	}
	defer db.Close()
	schema.DropSchema(db) // método temporário para realizar testes
	schema.CreateSchema(db)

	userController := &controllers.UserController{}
	userController.Init(db)
	chatController := &controllers.ChatController{}
	chatController.Init(db)

	homeRoutes()
	userRoutes(userController)
	chatRoutes(chatController)
}

// Run roda o servidor com ele escutando na porta 8080
func Run() {
	router.Run(":8080")
}

func homeRoutes() {
	router.GET("/", controllers.GetHomePage)
}

func userRoutes(userController *controllers.UserController) {
	router.POST("/usuarios", userController.PostUsers)
	router.GET("/usuarios", userController.GetAllUsers)
	router.GET("/usuarios/:userID", userController.GetUser)
}

func chatRoutes(chatController *controllers.ChatController) {
	router.POST("/chat", chatController.PostChat)
	router.POST("/chat/:chatID", chatController.PostInChat)
	router.GET("/chat/:chatID", chatController.GetChat)
}
