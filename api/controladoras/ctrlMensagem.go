package controladoras

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/repositorios"
)

func InserirMensagem(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err == nil {
		obj := struct {
			Conteudo string
			AutorID  int
		}{}
		err := c.BindJSON(&obj)
		fmt.Println("Obj: ", obj)
		if err != nil {
			c.JSON(http.StatusNoContent, err.Error())
		} else {
			repositorios.SetMensagem(obj.Conteudo, chatID, obj.AutorID)
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func RecuperarMensagens(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err == nil {
		msgs, err := repositorios.GetChatMensagens(chatID)
		if err == nil {
			c.JSON(http.StatusOK, msgs)
		} else {
			c.JSON(http.StatusNotFound, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
