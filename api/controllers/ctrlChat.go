package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/repositories"
)

// ChatController struct contendo seus repositórios
type ChatController struct {
	chatRepository *repositories.RepChat
	userRepository *repositories.RepUser
}

// Init inicializa os repositórios
func (chatctrl *ChatController) Init(db *sql.DB) {
	chatctrl.chatRepository = &repositories.RepChat{}
	chatctrl.chatRepository.Init(db)
	chatctrl.userRepository = &repositories.RepUser{}
	chatctrl.userRepository.Init(db)
}

// InserirChat envia o chat ao front
func (chatctrl *ChatController) InserirChat(c *gin.Context) {
	var err error
	obj := struct {
		Nome       string
		criadorID  int
		FotoPerfil string
	}{}
	err = c.BindJSON(&obj)
	if err != nil {
		c.JSON(http.StatusNoContent, err.Error())
	} else {
		chatID, err := chatctrl.chatRepository.SetChat(obj.Nome, obj.criadorID)
		if err != nil {
			c.JSON(http.StatusNoContent, err.Error())
		} else {
			err = chatctrl.chatRepository.SetUserInChat(chatID, obj.criadorID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}

// RecuperarChat manda os chats para o front
func (chatctrl *ChatController) RecuperarChat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("chatID"))
	if err == nil {
		ch, err := chatctrl.chatRepository.GetChat(id)
		if err == nil {
			c.JSON(http.StatusOK, ch)
		} else {
			c.JSON(http.StatusNotFound, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

// RecuperarTodosChats manda todos os chats do bd para o front
func (chatctrl *ChatController) RecuperarTodosChats(c *gin.Context) {
	chs, err := chatctrl.chatRepository.GetAllChats()
	if err == nil {
		c.JSON(http.StatusOK, chs)
	} else {
		c.JSON(http.StatusNotFound, err.Error())
	}
}
