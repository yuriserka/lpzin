package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/repositories"
)

type MessageController struct {
	messageRepository *repositories.RepMessage
	userRepository    *repositories.RepUser
	chatRepository    *repositories.RepChat
}

func (msgctrl *MessageController) Init(db *sql.DB) {
	msgctrl.messageRepository = &repositories.RepMessage{}
	msgctrl.messageRepository.Init(db)
	msgctrl.userRepository = &repositories.RepUser{}
	msgctrl.userRepository.Init(db)
	msgctrl.chatRepository = &repositories.RepChat{}
	msgctrl.chatRepository.Init(db)
}

func (msgctrl *MessageController) InserirMensagem(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err == nil {
		obj := struct {
			Conteudo string
			AutorID  int
		}{}
		err := c.BindJSON(&obj)
		if err != nil {
			c.JSON(http.StatusNoContent, err.Error())
		} else {
			msgctrl.messageRepository.SetMsg(obj.AutorID, chatID, obj.Conteudo)
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (msgctrl *MessageController) RecuperarMensagens(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err == nil {
		msgs, err := msgctrl.chatRepository.GetChatMsgs(chatID)
		if err == nil {
			c.JSON(http.StatusOK, msgs)
		} else {
			c.JSON(http.StatusNotFound, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
