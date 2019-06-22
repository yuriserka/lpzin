package repositorios

import (
	"database/sql"

	"github.com/yuriserka/lpzin/api/models"
)

// SetChat insere um chat no banco de dados
func SetChat(nome string) (int, error) {
	sqlString := `
		insert into chat (nome) values($1)
			returning id
	`
	var id int
	err := db.QueryRow(sqlString, nome).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetChat devolve um chat no banco de dados
func GetChat(chatID int) (*models.Chat, error) {
	sqlString := `
		select nome from chat where id = $1
	`
	var nome string
	row := db.QueryRow(sqlString, chatID)
	err := row.Scan(&nome)
	if err != nil {
		return nil, err
	}
	usuarios, err := GetUsuariosChat(chatID)
	if err != nil {
		return nil, err
	}
	mensagens, err := GetMensagensChat(chatID)
	if err != nil {
		return nil, err
	}
	return &models.Chat{
		ID:         chatID,
		Nome:       nome,
		Mensagens:  mensagens,
		Users:      usuarios,
		FotoPerfil: "",
	}, nil
}

func GetTodosChats() ([]*models.Chat, error) {
	sqlString := `
		select id from chat
	`
	rows, err := db.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats, err := getChatsFromRows(rows)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func AddChatMembro(chatID, userID int) error {
	sqlString := `
		insert into chat_tem_usuario (idchat, idusuario)
			values($1, $2)
	`
	_, err := db.Exec(sqlString, chatID, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetUsuariosChat(chatID int) ([]*models.Usuario, error) {
	sqlString := `
		select u.id from chat_tem_usuario as chu join usuario as u on
			chu.idusuario = u.id where chu.idchat = $1
	`
	rows, err := db.Query(sqlString, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios, err := getUsersFromRows(rows)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func GetMensagensChat(chatID int) ([]*models.Mensagem, error) {
	sqlString := `
		select idmsg from mensagem where idchat = $1
	`
	rows, err := db.Query(sqlString, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	mensagens, err := getMsgsFromRows(rows)
	if err != nil {
		return nil, err
	}

	return mensagens, nil
}

func getChatsFromRows(rows *sql.Rows) ([]*models.Chat, error) {
	var chats []*models.Chat
	var chatid int
	for rows.Next() {
		err := rows.Scan(&chatid)
		if err != nil {
			return nil, err
		}
		chat, err := GetChat(chatid)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
