package controladoras

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/repositorios"
)

func RecuperarChat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("chatID"))
	if err == nil {
		ch, err := repositorios.GetChat(id)
		if err == nil {
			c.JSON(http.StatusOK, ch)
		} else {
			c.JSON(http.StatusNotFound, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func InserirChat(c *gin.Context) {
	obj := struct {
		Nome       string
		CriadorID  int
		FotoPerfil string
	}{}
	err := c.BindJSON(&obj)
	if err != nil {
		c.JSON(http.StatusNoContent, err.Error())
	} else {
		chatID, err := repositorios.SetChat(obj.Nome)
		if err != nil {
			c.JSON(http.StatusNoContent, err.Error())
		} else {
			repositorios.AddChatMembro(chatID, obj.CriadorID)
		}
	}
}

func RecuperarTodosChats(c *gin.Context) {
	chs, err := repositorios.GetTodosChats()
	if err == nil {
		c.JSON(http.StatusOK, chs)
	} else {
		c.JSON(http.StatusNotFound, err.Error())
	}
}
