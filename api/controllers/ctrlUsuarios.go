package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/yuriserka/lpzin/api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/models"
)

type UserController struct {
	userRepository *repositories.RepUser
}

func (user *UserController) Init(db *sql.DB) {
	user.userRepository = &repositories.RepUser{}
	user.userRepository.Init(db)
}

func (user *UserController) PostUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"usuarios": models.Usuarios,
	})
}

func (user *UserController) GetAllUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, models.Usuarios)
}

func (user *UserController) GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if userID, err := strconv.Atoi(c.Param("userID")); err == nil {
		for _, v := range models.Usuarios {
			if v.ID == userID {
				c.JSON(http.StatusOK, &v)
				break
			}
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
