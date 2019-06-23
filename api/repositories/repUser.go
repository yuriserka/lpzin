package repositories

import (
	"database/sql"
	"errors"

	"github.com/yuriserka/lpzin/api/models"
	"golang.org/x/crypto/bcrypt"
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
func (rep *RepUser) SetUser(nome, foto, username, senha string) (int, error) {
	if len(nome) > 100 {
		return -1, errors.New("O nome deve possuir até 100 caracteres")
	}
	if len(username) > 100 {
		return -1, errors.New("O nome de usuário deve possuir até 100 caracteres")
	}

	var (
		hash []byte
		err  error
	)

	hash, err = encryptPass([]byte(senha))
	if err != nil {
		return -1, err
	}

	sqlStatement := `
	INSERT INTO Usuario (nome, foto, username, senha)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := -1

	err = rep.db.QueryRow(sqlStatement, nome, foto, username, string(hash)).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetUser retorna um usuário de acordo com a ID passada, sem sua senha
func (rep *RepUser) GetUser(userid int) (*models.Usuario, error) {
	var (
		id        int
		nome      string
		foto      string
		username  string
		ultimaVez string
	)
	sqlStatement := `SELECT id, nome, foto, username, ultima_vez FROM Usuario WHERE id = $1`
	value := rep.db.QueryRow(sqlStatement, userid)

	switch err := value.Scan(&id, &nome, &foto, &username, &ultimaVez); err {
	case sql.ErrNoRows:
		return nil, errors.New("Usuário não encontrado")
	case nil:
		break
	default:
		return nil, err
	}

	user := &models.Usuario{ID: id, Nome: nome, FotoPerfil: foto, Username: username, UltimaVez: ultimaVez}

	return user, nil
}

func (rep *RepUser) GetAllUsers() ([]*models.Usuario, error) {
	sqlStatement := `
	select id from usuario
	`
	rows, err := rep.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios, err := getUsersFromRows(rows, rep.db)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

// GetUserID recebe o username do usuário e retorna o id
func (rep *RepUser) GetUserID(username string) (int, error) {
	id := -1
	sqlStatement := `
	SELECT id FROM Usuario WHERE username = $1
	`
	value := rep.db.QueryRow(sqlStatement, username)

	switch err := value.Scan(&id); err {
	case sql.ErrNoRows:
		return -1, errors.New("Usuário não encontrado")
	case nil:
		break
	default:
		return -1, err
	}

	return id, nil
}

// GetUserMsgs retorna todas as mensagens de um usuário
func (rep *RepUser) GetUserMsgs(userid int) ([]*models.Mensagem, error) {
	sqlStatement := `
	SELECT IDMsg, Hr_env, Msg, IDChat, IDUsuario
	FROM Mensagem WHERE IDUsuario = $1
	`
	rows, err := rep.db.Query(sqlStatement, userid)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("Mensagens não encontradas")
	case nil:
		break
	default:
		return nil, err
	}
	msgs, err := getUserMsgsFromRows(rows, rep.db)
	if err != nil {
		return nil, err
	}

	e1 := rows.Close()
	if e1 != nil {
		return nil, e1
	}

	return msgs, nil
}

func (rep *RepUser) GetUserChats(userid int) ([]*models.Chat, error) {
	sqlStatement := `
	SELECT idchat FROM chat_tem_usuario WHERE IDUsuario = $1
	`
	rows, err := rep.db.Query(sqlStatement, userid)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("Chats não encontrados")
	case nil:
	default:
		return nil, err
	}
	defer rows.Close()

	chats, err := getChatsFromRows(rows, rep.db)

	return chats, nil
}

// UserAuth recebe o username e senha, retorna true se a senha for correta e false caso contrário
func (rep *RepUser) UserAuth(username, senha string) (bool, error) {
	auth, err := validatePass(rep.db, username, senha)
	if err != nil {
		return false, err
	}
	return auth, nil
}

func getUserMsgsFromRows(rows *sql.Rows, db *sql.DB) ([]*models.Mensagem, error) {
	var msgs []*models.Mensagem
	var (
		id, chatID, autorID int
		hrEnv, conteudo     string
	)
	for rows.Next() {
		err := rows.Scan(&id, &hrEnv, &conteudo, &chatID, &autorID)
		if err != nil {
			return nil, err
		}
		msg := &models.Mensagem{
			ID:        id,
			HoraEnvio: hrEnv,
			Conteudo:  conteudo,
			ChatID:    chatID,
			Autor:     autorID,
		}
		msgs = append(msgs, msg)
	}
	if msgs == nil {
		return nil, errors.New("Usuário não possui mensagens enviadas")
	}
	return msgs, nil
}

func getUsersFromRows(rows *sql.Rows, db *sql.DB) ([]*models.Usuario, error) {
	var users []*models.Usuario
	rep := RepUser{}
	rep.Init(db)
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		user, err := rep.GetUser(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if users == nil {
		return nil, errors.New("Não possui usuários")
	}
	return users, nil
}

func encryptPass(senha []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(senha, bcrypt.DefaultCost)
	return hash, err
}

func getHashPassword(db *sql.DB, username string) (string, error) {
	sqlStatement := `
	SELECT senha FROM Usuario
	WHERE username = $1
	`
	var hash string

	err := db.QueryRow(sqlStatement, username).Scan(&hash)
	return hash, err
}

func validatePass(db *sql.DB, username, senha string) (bool, error) {
	var (
		pass     []byte
		hashpass string
		err      error
	)

	pass = []byte(senha)
	hashpass, err = getHashPassword(db, username)
	if err != nil {
		return false, errors.New("Usuário não existente")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashpass), pass)
	if err != nil {
		return false, err
	}

	return true, nil
}
