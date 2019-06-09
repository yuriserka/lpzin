package routers

import (
	"fmt"
	"strconv"

	"github.com/yuriserka/lpzin/api/repositories"

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

	r := repositories.RepUser{}
	r.Init(db)
	id := r.SetUser("Yuri Serka", "foto de umgadokkk")
	user := r.GetUser(strconv.Itoa(id))
	fmt.Println(user)

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
