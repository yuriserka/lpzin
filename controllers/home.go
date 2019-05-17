package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/models"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	usuario := models.Yuri
	render(c, gin.H{
		"title":   "Home",
		"payload": articles,
		"user":    usuario,
	}, "index.html")
}
