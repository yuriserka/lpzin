package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/models"
)

func PostChat(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"chats": models.Chats,
	})
}

func PostInChat(c *gin.Context) {
	var message models.Mensagem
	c.BindJSON(&message)

	for i := 0; i < len(models.Chats); i++ {
		chat := &models.Chats[i]
		if chat.ID == message.ChatID {
			chat.Mensagens = append(chat.Mensagens, message)
			break
		}
	}
}

func GetChat(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if chatID, err := strconv.ParseUint(c.Param("chatID"), 10, 64); err == nil {
		for _, val := range models.Chats {
			if val.ID == chatID {
				c.JSON(http.StatusOK, &val.Mensagens)
				break
			}
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
