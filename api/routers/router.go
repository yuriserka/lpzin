package routers

import (
	"github.com/yuriserka/lpzin/api/controladoras"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	// db, err = common.ConnDB()
	// if err != nil {
	// 	log.Panic(fmt.Sprintf("db: %v", err))
	// }
	userRoutes()
	chatRoutes()

	// schema.DropSchema(db) // método temporário para realizar testes
	// schema.CreateSchema(db)
}

// Run roda o servidor com ele escutando na porta 8080
func Run() {

	// userController := &controllers.UserController{}
	// userController.Init(db)
	// chatController := &controllers.ChatController{}
	// chatController.Init(db)
	// homeRoutes()
	// userRoutes(userController)
	// chatRoutes(chatController)
	// defer db.Close()

	router.Run(":8080")
}

func userRoutes( /*userController *controllers.UserController*/ ) {
	router.GET("/usuarios", controladoras.RecuperarTodosUsuarios)
	router.GET("/usuarios/:userID", controladoras.RecuperarUsuario)
	router.POST("/usuarios", controladoras.InserirUsuario)
	router.GET("/usuarios/:userID/chats", controladoras.RecuperarChatsParticipantes)
}

func chatRoutes( /*chatController *controllers.ChatController*/ ) {
	router.GET("/chats", controladoras.RecuperarTodosChats)
	router.GET("/chats/:chatID", controladoras.RecuperarChat)
	router.POST("/chats", controladoras.InserirChat)
	router.POST("/chats/:chatID/mensagens", controladoras.InserirMensagem)
	router.GET("/chats/:chatID/mensagens", controladoras.RecuperarMensagens)
}
