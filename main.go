package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	ID   int    `json:"id"`
	Nick string `json:"username"`
	Msg  string `json:"message"`
}

func main() {
	db := make(map[int]usuario)
	contador := 0

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./lpzin_frontend/build", true)))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/teste", func(c *gin.Context) {
		usr := usuario{
			ID:   contador,
			Nick: c.PostForm("nick"),
			Msg:  c.PostForm("message"),
		}
		db[usr.ID] = usr
		contador++
		c.JSON(http.StatusOK, gin.H{
			"usuario": usr,
		})
	})

	router.GET("/teste", func(c *gin.Context) {
		for _, v := range db {
			c.Data(http.StatusOK, "text", []byte(v.Nick+"\n"))
		}
	})

	router.Run(":8080")
}
