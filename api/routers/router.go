package routers

import (
	"fmt"

	"github.com/yuriserka/lpzin/api/common"
	"github.com/yuriserka/lpzin/api/controllers"
	"github.com/yuriserka/lpzin/api/repositories"
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
		panic(fmt.Sprintf("db: %v", err))
	}
	defer db.Close()
	schema.DropSchema(db) // método temporário para realizar testes
	schema.CreateSchema(db)

	// Inicialização dos repositórios a fim de teste, depois mudar as estruturas dos controllers para ter que só inicializar o controller aqui
	ruser := repositories.RepUser{}
	ruser.Init(db)
	rchat := repositories.RepChat{}
	rchat.Init(db)

	homeRoutes()
	userRoutes()
	chatRoutes()
}

// Run roda o servidor com ele escutando na porta 8080
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
}

func chatRoutes() {
	router.POST("/chat", controllers.PostChat)
	router.POST("/chat/:chatID", controllers.PostInChat)
	router.GET("/chat/:chatID", controllers.GetChat)
}
