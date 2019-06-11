package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yuriserka/lpzin/api/models"
)

// RepMessage contém a instância do BD a ser utilizada
type RepMessage struct {
	db *sql.DB
}

// Init recebe a instância do banco de dados e inicializa na struct
func (rep *RepMessage) Init(db *sql.DB) {
	rep.db = db
}

// SetMsg cria uma mensagem indicando qual o id do usuário, chat e o conteúdo da msg e retorna o id da msg criada
func (rep *RepMessage) SetMsg(userid, chatid, msg string) int {
	if len(msg) > 500 {
		log.Panic("A mensagem deve conter até 500 caracteres")
	}
	id := 0
	sqlStatement := `
	INSERT INTO Mensagem (Msg, IDChat, IDUsuario)
	VALUES ($1, $2, $3)
	RETURNING IDMsg`

	fmt.Println("msg no setmsg:", msg)

	err := rep.db.QueryRow(sqlStatement, msg, chatid, userid).Scan(&id)
	if err != nil {
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}

	fmt.Println("ID da mensagem criada: ", id)

	return id
}

// GetMsg passado a id da mensagem é retornado uma mensagem
func (rep *RepMessage) GetMsg(msgid string) models.Mensagem {
	msgret := models.Mensagem{}
	sqlStatement := `
	SELECT IDMsg, Hr_env, Msg, IDChat, IDUsuario
	FROM Mensagem WHERE IDMsg = $1`

	value := rep.db.QueryRow(sqlStatement, msgid)

	switch err := value.Scan(&msgret.ID, &msgret.HoraEnvio, &msgret.Conteudo, &msgret.ChatID, &msgret.Autor); err {
	case sql.ErrNoRows:
		fmt.Println("Mensagem não encontrada")
	case nil:
	default:
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}

	return msgret
}
