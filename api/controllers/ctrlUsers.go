package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/yuriserka/lpzin/api/repositories"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository *repositories.RepUser
}

func (userctrl *UserController) Init(db *sql.DB) {
	userctrl.userRepository = &repositories.RepUser{}
	userctrl.userRepository.Init(db)
}

func (userctrl *UserController) InserirUsuario(c *gin.Context) {
	obj := struct {
		Nome       string
		FotoPerfil string
		Username   string
		UltimaVez  string
		Senha      string
	}{}
	err := c.BindJSON(obj)
	if err != nil {
		c.JSON(http.StatusNoContent, err.Error())
	} else {
		_, err := userctrl.userRepository.SetUser(obj.Nome, obj.FotoPerfil, obj.Username, obj.Senha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

func (userctrl *UserController) RecuperarChatsParticipantes(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err == nil {
		chats, err := userctrl.userRepository.GetUserChats(userID)
		if err == nil {
			c.JSON(http.StatusOK, chats)
		} else {
			c.JSON(http.StatusNotFound, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (userctrl *UserController) RecuperarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userID"))
	if err == nil {
		u, err := userctrl.userRepository.GetUser(id)
		if err == nil {
			c.JSON(http.StatusOK, u)
		} else {
			c.JSON(http.StatusNotFound, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (userctrl *UserController) RecuperarTodosUsuarios(c *gin.Context) {
	us, err := userctrl.userRepository.GetAllUsers()
	if err == nil {
		c.JSON(http.StatusOK, us)
	} else {
		c.JSON(http.StatusNotFound, err.Error())
	}
}
