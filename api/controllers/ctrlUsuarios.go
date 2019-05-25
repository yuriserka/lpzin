package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/models"
)

func PostUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"usuarios": models.Usuarios,
	})
}

func GetAllUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, models.Usuarios)
}

func GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if userID, err := strconv.Atoi(c.Param("userID")); err == nil {
		for _, v := range models.Usuarios {
			if v.ID == userID {
				c.JSON(http.StatusOK, &v)
			}
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
