package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuriserka/lpzin/api/models"
	"github.com/yuriserka/lpzin/api/repositories"
)

// ChatController struct contendo seus repositórios
type ChatController struct {
	chatRepository *repositories.RepChat
	userRepository *repositories.RepUser
}

// Init inicializa os repositórios
func (chat *ChatController) Init(db *sql.DB) {
	chat.chatRepository = &repositories.RepChat{}
	chat.chatRepository.Init(db)

	chat.userRepository = &repositories.RepUser{}
	chat.userRepository.Init(db)
}

// PostChat envia os chats
func (chat *ChatController) PostChat(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"chats": models.Chats,
	})
}

// PostInChat envia uma mensagem a um chat
func (chat *ChatController) PostInChat(c *gin.Context) {
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

// GetChat manda os chats
func (chat *ChatController) GetChat(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if chatID, err := strconv.Atoi(c.Param("chatID")); err == nil {
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
