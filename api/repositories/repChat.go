package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/yuriserka/lpzin/api/models"
)

// RepChat contém a instância do BD a ser utilizada
type RepChat struct {
	db *sql.DB
}

// Init recebe a instância do banco de dados e inicializa na struct
func (rep *RepChat) Init(db *sql.DB) {
	rep.db = db
}

// SetChat recebe o nome do chat e o usuário criador do chat e retorna o id do chat
func (rep *RepChat) SetChat(nome, userid string) int {
	if len(nome) > 100 {
		log.Panic("O nome do chat deve possuir até 100 caracteres")
	}
	id := 0
	sqlStatement := `
	INSERT INTO Chat (nome) 
	VALUES ($1) 
	RETURNING id`
	err := rep.db.QueryRow(sqlStatement, nome).Scan(&id)
	if err != nil {
		log.Printf("Error %+v\n", err)
	}
	// Adicionando o usuario ao chat
	rep.SetChatUser(strconv.Itoa(id), userid)

	fmt.Println("ID do Chat criado: ", id)

	return id
}

func (rep *RepChat) SetChatUser(chatid, userid string) {
	sqlStatement := `
	INSERT INTO Chat_tem_usuario(idchat, idusuario)
	VALUES ($1, $2)`

	_, err := rep.db.Exec(sqlStatement, chatid, userid)
	if err != nil {
		log.Printf("Error %+v\n", err)
	}
}

func (rep *RepChat) GetChat(chatid string) models.Chat {
	var (
		id        int
		nome      string
		users     []models.Usuario
		mensagens []models.Mensagem
	)
	sqlStatement := `
	SELECT id, nome FROM Chat WHERE id = $1
	`
	value := rep.db.QueryRow(sqlStatement, chatid)

	switch err := value.Scan(&id, &nome); err {
	case sql.ErrNoRows:
		fmt.Println("Chat não encontrado")
	case nil:
	default:
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}

	users = rep.getChatUsers(strconv.Itoa(id))

	// Fazer getChatMsgs

	chat := models.Chat{ID: id, Nome: nome, Users: users, Mensagens: mensagens}

	return chat
}

func (rep *RepChat) getChatUsers(chatid string) []models.Usuario {
	sqlStatement := `
	SELECT u.id FROM Usuario u
	JOIN Chat_tem_usuario chu ON u.id = chu.idusuario
	JOIN Chat c ON chu.idchat = c.id
	WHERE c.id = $1
	`
	rows, err := rep.db.Query(sqlStatement, chatid)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Chat não encontrado")
	case nil:
	default:
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}
	defer rows.Close()

	users, err := getUsersFromRows(rows, rep.db)
	if err != nil {
		log.Panic(fmt.Sprintf("Error %+v\n", err))
	}

	return users
}

func getUsersFromRows(rows *sql.Rows, db *sql.DB) ([]models.Usuario, error) {
	users := []models.Usuario{}
	rep := RepUser{}
	rep.Init(db)
	for rows.Next() {
		var (
			userid int
			user   models.Usuario
		)
		err := rows.Scan(&userid)
		if err != nil {
			return nil, err
		}
		user = rep.GetUser(strconv.Itoa(userid))
		users = append(users, user)
	}
	return users, nil
}
