package repositories

import (
	"database/sql"
	"errors"

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
func (rep *RepMessage) SetMsg(userID, chatID int, msg string) (int, error) {
	var (
		id  int
		err error
	)

	if len(msg) > 500 {
		err = errors.New("A mensagem deve conter até 500 caracteres")
		return -1, err
	}

	sqlStatement := `
	INSERT INTO Mensagem (Msg, IDChat, IDUsuario)
	VALUES ($1, $2, $3)
	RETURNING IDMsg`

	err = rep.db.QueryRow(sqlStatement, msg, chatID, userID).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetMsg passado a id da mensagem é retornado uma mensagem
func (rep *RepMessage) GetMsg(msgID int) (*models.Mensagem, error) {
	var err error
	msgret := &models.Mensagem{}
	sqlStatement := `
	SELECT IDMsg, Hr_env, Msg, IDChat, IDUsuario
	FROM Mensagem WHERE IDMsg = $1`

	value := rep.db.QueryRow(sqlStatement, msgID)

	switch err = value.Scan(&msgret.ID, &msgret.HoraEnvio, &msgret.Conteudo, &msgret.ChatID, &msgret.Autor); err {
	case sql.ErrNoRows:
		err = errors.New("Mensagem não encontrada")
		return nil, err
	case nil:
		break
	default:
		return nil, err
	}

	return msgret, nil
}
