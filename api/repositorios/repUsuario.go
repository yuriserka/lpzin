package repositorios

import (
	"database/sql"

	"github.com/yuriserka/lpzin/api/models"
	"github.com/yuriserka/lpzin/api/utils"
)

// GetUsuario retorna um usuario fo banco de dados, mas sem sua senha
func GetUsuario(userID int) (*models.Usuario, error) {
	sqlString := `
		select nome, foto, username, ultima_vez from usuario where id = $1 
	`
	var (
		nome, foto, username, visto string
	)
	row := db.QueryRow(sqlString, &userID)
	err := row.Scan(&nome, &foto, &username, &visto)

	if err != nil {
		return nil, err
	}

	return &models.Usuario{
		ID:         userID,
		Nome:       nome,
		FotoPerfil: foto,
		Username:   username,
		UltimaVez:  visto,
	}, nil
}

// SetUsuario armazena um usuario no banco de dados
func SetUsuario(nome, foto, username, senha string) (int, error) {
	sqlString := `
		insert into usuario
			(nome, foto, username, senha) values($1, $2, $3, $4)
		returning id
	`
	var id int
	senhaCriptografada := utils.CriptografaSenha([]byte(senha))
	err := db.QueryRow(sqlString, nome, foto, username,
		senhaCriptografada).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetSenhaUsuario retorna s senha do usuario
func GetSenhaUsuario(userID int) (string, error) {
	sqlString := `
		select senha from usuario where $1
	`
	var senha string
	row := db.QueryRow(sqlString, userID)
	err := row.Scan(&senha)
	if err != nil {
		return "", err
	}
	return senha, nil
}

func GetUserChats(userid int) ([]*models.Chat, error) {
	sqlStatement := `
		select idchat from chat_tem_usuario where idusuario = $1
	`
	rows, err := db.Query(sqlStatement, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chatsIDs, err := getChatsFromRows(rows)

	return chatsIDs, err
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
