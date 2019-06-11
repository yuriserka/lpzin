package repositories

import(
	"database/sql"
	"fmt"
	"log"
	"github.com/yuriserka/lpzin/models"
)

type RepMessage struct {
	db *sql.DB
}

func (rep *RepMessage) Init(db *sql.DB) {
	rep.db = db
}

func (rep *RepMessage) SetMsg(userid, chatid, msg string) int {
	if len(msg) > 500 {
		log.Panic("A mensagem deve conter até 500 caracteres")
	}
	id := 0
	sqlStatement := `
	INSERT INTO Mensagem (Msg, IDChat, IDUsuario)
	VALUES ($1, $2, $3)
	RETURNING IDMsg`

	_, err := rep.db.Exec(sqlStatement, msg, chatid, userid)
	if err != nil {
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}

	fmt.Println("ID da mensagem criada: ", id)

	return id
}

func (rep *RepMessage) GetMsg(msgid string) models.Mensagem {
	msg := models.Mensagem{}

	return msg
}