package controladoras

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/yuriserka/lpzin/api/repositorios"
// )

// func RecuperarUsuario(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("userID"))
// 	if err == nil {
// 		u, err := repositorios.GetUsuario(id)
// 		if err == nil {
// 			c.JSON(http.StatusOK, u)
// 		} else {
// 			c.JSON(http.StatusNotFound, err.Error())
// 		}
// 	} else {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 	}
// }

// func InserirUsuario(c *gin.Context) {
// 	obj := struct {
// 		Nome       string
// 		FotoPerfil string
// 		Username   string
// 		UltimaVez  string
// 		Senha      string
// 	}{}
// 	err := c.BindJSON(obj)
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, err.Error())
// 	} else {
// 		repositorios.SetUsuario(obj.Nome, obj.FotoPerfil, obj.Username, obj.Senha)
// 	}
// }

// func RecuperarChatsParticipantes(c *gin.Context) {
// 	userID, err := strconv.Atoi(c.Param("userID"))
// 	if err == nil {
// 		chats, err := repositorios.GetChatsUsuario(userID)
// 		if err == nil {
// 			c.JSON(http.StatusOK, chats)
// 		} else {
// 			c.JSON(http.StatusNotFound, err.Error())
// 		}
// 	} else {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 	}
// }

// func RecuperarTodosUsuarios(c *gin.Context) {
// 	us, err :=repositorios.GetTodosUsuarios()
// 	if err == nil {
// 		c.JSON(http.StatusOK, us)
// 	} else {
// 		c.JSON(http.StatusNotFound, err.Error())
// 	}
// }
