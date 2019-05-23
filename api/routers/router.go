package routers

import (
	"net/http"

	"github.com/yuriserka/lpzin/api/models"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/teste", func(c *gin.Context) {
		usr := models.Usuario{
			ID:   models.Contador,
			Nick: c.PostForm("nick"),
			Msg:  c.PostForm("message"),
		}
		models.Db[usr.ID] = usr
		models.Contador++
		c.JSON(http.StatusOK, gin.H{
			"usuario": usr,
		})
	})

	router.GET("/teste", func(c *gin.Context) {
		for _, v := range models.Db {
			c.Data(http.StatusOK, "text", []byte(v.Nick+"\n"))
		}
	})
}

func Run() {
	router.Run(":8080")
}
