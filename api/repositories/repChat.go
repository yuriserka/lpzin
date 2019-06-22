package repositories

import (
	"database/sql"
	"errors"

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
func (rep *RepChat) SetChat(nome string, userID int) (int, error) {
	var (
		id  int
		err error
	)

	if len(nome) > 100 {
		err = errors.New("O nome do chat deve possuir até 100 caracteres")
		return -1, err
	}

	sqlStatement := `
	INSERT INTO Chat (nome)
	VALUES ($1)
	RETURNING id`
	err = rep.db.QueryRow(sqlStatement, nome).Scan(&id)
	if err != nil {
		return -1, err
	}
	// Adicionando o usuario ao chat
	err = rep.SetUserInChat(id, userID)

	return id, err
}

// SetUserInChat seta o usuário pasado como participante do chat passado
func (rep *RepChat) SetUserInChat(chatID, userID int) error {
	sqlStatement := `
	INSERT INTO Chat_tem_usuario(idchat, idusuario)
	VALUES ($1, $2)`

	_, err := rep.db.Exec(sqlStatement, chatID, userID)
	return err
}

// GetChat retorna o chat de acordo com o id do chat passado
func (rep *RepChat) GetChat(chatID int) (*models.Chat, error) {
	var (
		id        int
		nome      string
		users     []*models.Usuario
		mensagens []*models.Mensagem
		err       error
	)
	sqlStatement := `
	SELECT id, nome FROM Chat WHERE id = $1
	`
	value := rep.db.QueryRow(sqlStatement, chatID)

	switch err = value.Scan(&id, &nome); err {
	case sql.ErrNoRows:
		err = errors.New("Chat não encontrado")
		return nil, err
	case nil:
		break
	default:

	}

	users, err = rep.GetChatUsers(id)
	if err != nil {
		return nil, err
	}

	mensagens, err = rep.GetChatMsgs(id)
	if err != nil {
		return nil, err
	}

	chat := &models.Chat{ID: id, Nome: nome, Users: users, Mensagens: mensagens}

	return chat, err
}

//GetAllChats retorna todos os chats armazenados no banco de dados
func (rep *RepChat) GetAllChats() ([]*models.Chat, error) {
	sqlStatement := `
	SELECT Id from Chat
	`
	rows, err := rep.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats, err := getChatsFromRows(rows, rep.db)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// GetChatUsers recebe o id do chat e retorna quais são os usuários que fazem parte do chat
func (rep *RepChat) GetChatUsers(chatID int) ([]*models.Usuario, error) {
	var (
		rows  *sql.Rows
		users []*models.Usuario
		err   error
	)
	sqlStatement := `
	SELECT u.id FROM Usuario u
	JOIN Chat_tem_usuario chu ON u.id = chu.idusuario
	JOIN Chat c ON chu.idchat = c.id
	WHERE c.id = $1
	`
	rows, err = rep.db.Query(sqlStatement, chatID)
	switch err {
	case sql.ErrNoRows:
		err = errors.New("Chat não encontrado")
		return nil, err
	case nil:
		break
	default:
		return nil, err
	}
	defer rows.Close()

	users, err = getUsersFromRows(rows, rep.db)
	if err != nil {
		return nil, err
	}

	return users, err
}

// GetChatMsgs recupera as mensagens de um chat de acordo com o id do chat passado
func (rep *RepChat) GetChatMsgs(chatID int) ([]*models.Mensagem, error) {
	msgs := []*models.Mensagem{}
	sqlStatement := `
	SELECT IDMsg FROM Mensagem
	WHERE IDChat = $1;
	`
	rows, err := rep.db.Query(sqlStatement, chatID)
	switch err {
	case sql.ErrNoRows:
		err := errors.New("O chat não possui mensagens")
		return nil, err
	case nil:
		break
	default:
		return nil, err
	}
	defer rows.Close()

	msgs, err = getMsgsFromRows(rows, rep.db)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func getChatsFromRows(rows *sql.Rows, db *sql.DB) ([]*models.Chat, error) {
	rep := &RepChat{}
	rep.Init(db)
	var chats []*models.Chat
	var chatid int
	for rows.Next() {
		err := rows.Scan(&chatid)
		if err != nil {
			return nil, err
		}
		chat, err := rep.GetChat(chatid)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func getMsgsFromRows(rows *sql.Rows, db *sql.DB) ([]*models.Mensagem, error) {
	msgs := []*models.Mensagem{}
	rep := &RepMessage{}
	rep.Init(db)
	for rows.Next() {
		var (
			msgid int
			msg   *models.Mensagem
		)
		err := rows.Scan(&msgid)
		if err != nil {
			return nil, err
		}
		msg, err = rep.GetMsg(msgid)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}
