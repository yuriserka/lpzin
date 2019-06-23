package routers

import (
	"database/sql"

	"github.com/yuriserka/lpzin/api/controllers"
	"github.com/yuriserka/lpzin/schema"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Init irá criar as rotas do aplicativo e também servirá a pasta estatica do front
func Init(db *sql.DB) {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	userController := &controllers.UserController{}
	userController.Init(db)
	chatController := &controllers.ChatController{}
	chatController.Init(db)
	msgController := &controllers.MessageController{}
	msgController.Init(db)

	userRoutes(userController)
	chatRoutes(chatController)
	messageRoutes(msgController)

	// schema.DropSchema(db) // método temporário para realizar testes
	schema.CreateSchema(db)
}

// func AuthenticationRequired(auths ... string) gin.HandlerFunc {
// 	session := sessions.Default(c)
// 	user := session.Get("user")
// 	if user == nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "user precisa estar cadastrado para usar os serviços"})
// 		c.Abort()
// 		return
// 	}

// 	if len(auths) != 0 {
// 		authType := session.Get("authType")
// 		if authType == nil || !funk.ContainsString(auths, authType.(string)) {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
// 			c.Abort()
// 			return
// 		}
// 	}
// }

// Run roda o servidor com ele escutando na porta 8080
func Run() {
	router.Run(":8080")
}

// func loginRoutes(userController *controllers.UserController) {
// 	//router.GET("/login", userController.login)
// 	router.POST("/login/:username:senha", userController.Login)

// 	router.GET("/cadastro")
// 	router.POST("/cadastro/:username:senha")
// }

func userRoutes(userController *controllers.UserController) {
	router.GET("/usuarios", userController.RecuperarTodosUsuarios)
	router.GET("/usuarios/:userID", userController.RecuperarUsuario)
	router.POST("/usuarios", userController.InserirUsuario)
	router.GET("/usuarios/:userID/chats", userController.RecuperarChatsParticipantes)
}

func chatRoutes(chatController *controllers.ChatController) {
	router.GET("/chats", chatController.RecuperarTodosChats)
	router.GET("/chats/:chatID", chatController.RecuperarChat)
	router.POST("/chats", chatController.InserirChat)
}

func messageRoutes(msgController *controllers.MessageController) {
	router.POST("/chats/:chatID/mensagens", msgController.InserirMensagem)
	router.GET("/chats/:chatID/mensagens", msgController.RecuperarMensagens)
}
