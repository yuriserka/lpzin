package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yuriserka/lpzin/api/models"
)

// RepUser contém a instância do BD a ser utilizada
type RepUser struct {
	db *sql.DB
}

// Init recebe a instância do banco de dados e inicializa na struct
func (rep *RepUser) Init(db *sql.DB) {
	rep.db = db
}

// SetUser cria um usuário no banco de dados e retorna o id do usuário criado
func (rep *RepUser) SetUser(nome string, foto string) int {
	if len(nome) > 100 {
		log.Panic("O nome do usuário deve possuir até 100 caracteres")
	}
	sqlStatement := `
	INSERT INTO Usuario (nome, foto)
	VALUES ($1, $2)
	RETURNING id`
	id := 0

	err := rep.db.QueryRow(sqlStatement, nome, foto).Scan(&id)
	if err != nil {
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}
	fmt.Println("ID do usuário criado: ", id)
	return id
}

// GetUser retorna um usuário de acordo com a ID passada
func (rep *RepUser) GetUser(userid int) models.Usuario {
	var (
		id   int
		nome string
		foto string
	)
	sqlStatement := `SELECT id, nome, foto FROM Usuario WHERE id = $1`
	value := rep.db.QueryRow(sqlStatement, userid)

	switch err := value.Scan(&id, &nome, &foto); err {
	case sql.ErrNoRows:
		fmt.Println("Usuário não encontrado")
	case nil:
	default:
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}

	user := models.Usuario{ID: id, Nome: nome, FotoPerfil: foto}

	return user
}

// GetUserMsgs retorna todas as mensagens de um usuário
func (rep *RepUser) GetUserMsgs(userid int) []models.Mensagem {
	sqlStatement := `
	SELECT IDMsg, Hr_env, Msg, IDChat, IDUsuario
	FROM Mensagem WHERE IDUsuario = $1
	`
	rows, err := rep.db.Query(sqlStatement, userid)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Mensagens não encontradas")
	case nil:
	default:
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}
	defer rows.Close()

	msgs, err := getUserMsgsFromRows(rows, rep.db)

	return msgs
}

func getUserMsgsFromRows(rows *sql.Rows, db *sql.DB) ([]models.Mensagem, error) {
	msgs := []models.Mensagem{}
	var msg models.Mensagem
	for rows.Next() {
		err := rows.Scan(&msg.ID, &msg.HoraEnvio, &msg.Conteudo, &msg.ChatID, &msg.Autor)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}
