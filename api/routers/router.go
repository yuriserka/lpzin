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

	db, err := utils.ConnDB()
	if err != nil {
		panic(fmt.Sprintf("db: %v", err))
	}
	defer db.Close()

	// Dar um jeito de fazer uma controladora para criar as tabelas antes da aplicação começar.
	// pode-se fazer até mesmo no ctrlHome se pá.

	// r := repositories.RepUser{}
	// r.Init(db)
	// id := r.SetUser("Yuri Serka", "foto de umgadokkk")
	// user := r.GetUser(strconv.Itoa(id))
	// fmt.Println(user)

	homeRoutes()
	userRoutes()
	chatRoutes()
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
}

func chatRoutes() {
	router.POST("/chat", controllers.PostChat)
	router.POST("/chat/:chatID", controllers.PostInChat)
	router.GET("/chat/:chatID", controllers.GetChat)
}
