package repositories

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	"github.com/pkg/errors"

// 	"github.com/yuriserka/lpzin/api/models"
// 	"golang.org/x/crypto/bcrypt"
// )

// // RepUser contém a instância do BD a ser utilizada
// type RepUser struct {
// 	db *sql.DB
// }

// // Init recebe a instância do banco de dados e inicializa na struct
// func (rep *RepUser) Init(db *sql.DB) {
// 	rep.db = db
// }

// // SetUser cria um usuário no banco de dados e retorna o id do usuário criado
// func (rep *RepUser) SetUser(nome, foto, username, senha string) int {
// 	if len(nome) > 100 {
// 		log.Panic("O nome deve possuir até 100 caracteres")
// 	}
// 	if len(username) > 100 {
// 		log.Panic("O nome de usuário deve possuir até 100 caracteres")
// 	}

// 	var (
// 		hash []byte
// 		err  error
// 	)

// 	hash, err = encryptPass([]byte(senha))
// 	if err != nil {
// 		log.Panic(fmt.Sprintf("Error %+v\n", err))
// 	}

// 	sqlStatement := `
// 	INSERT INTO Usuario (nome, foto, username, senha)
// 	VALUES ($1, $2, $3, $4)
// 	RETURNING id`
// 	id := 0

// 	err = rep.db.QueryRow(sqlStatement, nome, foto, username, string(hash)).Scan(&id)
// 	if err != nil {
// 		log.Panic(fmt.Sprintf("Error %+v\n", err))
// 	}
// 	fmt.Println("ID do usuário criado: ", id)
// 	return id
// }

// // GetUser retorna um usuário de acordo com a ID passada
// func (rep *RepUser) GetUser(userid int) (*models.Usuario, error) {
// 	var (
// 		id        int
// 		nome      string
// 		foto      string
// 		username  string
// 		ultimaVez string
// 	)
// 	sqlStatement := `SELECT id, nome, foto, username, ultima_vez FROM Usuario WHERE id = $1`
// 	value := rep.db.QueryRow(sqlStatement, userid)

// 	switch err := value.Scan(&id, &nome, &foto, &username, &ultimaVez); err {
// 	case sql.ErrNoRows:
// 		return nil, errors.New("Usuário não encontrado")
// 	case nil:
// 	default:
// 		log.Panic(fmt.Sprintf("Error %+v\n", err))
// 	}

// 	user := &models.Usuario{ID: id, Nome: nome, FotoPerfil: foto, Username: username, UltimaVez: ultimaVez}

// 	return user, nil
// }

// // GetUserID recebe o username do usuário e retorna o id
// func (rep *RepUser) GetUserID(username string) int {
// 	id := -1
// 	sqlStatement := `
// 	SELECT id FROM Usuario WHERE username = $1
// 	`
// 	value := rep.db.QueryRow(sqlStatement, username)

// 	switch err := value.Scan(&id); err {
// 	case sql.ErrNoRows:
// 		fmt.Println("Usuário não encontrado")
// 	case nil:
// 	default:
// 		log.Panic(fmt.Sprintf("Error %+v\n", err))
// 	}
// 	return id
// }

// // GetUserMsgs retorna todas as mensagens de um usuário
// func (rep *RepUser) GetUserMsgs(userid int) []models.Mensagem {
// 	sqlStatement := `
// 	SELECT IDMsg, Hr_env, Msg, IDChat, IDUsuario
// 	FROM Mensagem WHERE IDUsuario = $1
// 	`
// 	rows, err := rep.db.Query(sqlStatement, userid)
// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("Mensagens não encontradas")
// 	case nil:
// 	default:
// 		log.Panic(fmt.Sprintf("Error %+v\n", err))
// 	}
// 	defer rows.Close()

// 	msgs, err := getUserMsgsFromRows(rows, rep.db)

// 	return msgs
// }

// func (rep *RepUser) GetUserChats(userid int) []int {
// 	sqlStatement := `
// 	SELECT idchat FROM chat_tem_usuario WHERE IDUsuario = $1
// 	`
// 	rows, err := rep.db.Query(sqlStatement, userid)
// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("Chats não encontrados")
// 	case nil:
// 	default:
// 		log.Panic(fmt.Sprintf("Error %+v\n", err))
// 	}
// 	defer rows.Close()

// 	chatsIDs, err := getUserChatsIDsFromRows(rows, rep.db)

// 	return chatsIDs
// }

// // UserAuth recebe o username e senha, retorna true se a senha for correta e false caso contrário
// func (rep *RepUser) UserAuth(username, senha string) bool {
// 	auth := validatePass(rep.db, username, senha)
// 	return auth
// }

// func getUserMsgsFromRows(rows *sql.Rows, db *sql.DB) ([]models.Mensagem, error) {
// 	msgs := []models.Mensagem{}
// 	var msg models.Mensagem
// 	for rows.Next() {
// 		err := rows.Scan(&msg.ID, &msg.HoraEnvio, &msg.Conteudo, &msg.ChatID, &msg.Autor)
// 		if err != nil {
// 			return nil, err
// 		}
// 		msgs = append(msgs, msg)
// 	}
// 	return msgs, nil
// }

// func getUserChatsIDsFromRows(rows *sql.Rows, db *sql.DB) ([]int, error) {
// 	var chats []int
// 	var chatid int
// 	for rows.Next() {
// 		err := rows.Scan(&chatid)
// 		if err != nil {
// 			return nil, err
// 		}
// 		chats = append(chats, chatid)
// 	}
// 	return chats, nil
// }

// func encryptPass(senha []byte) ([]byte, error) {
// 	hash, err := bcrypt.GenerateFromPassword(senha, bcrypt.DefaultCost)
// 	return hash, err
// }

// func getHashPassword(db *sql.DB, username string) (string, error) {
// 	sqlStatement := `
// 	SELECT senha FROM Usuario
// 	WHERE username = $1
// 	`
// 	var hash string

// 	err := db.QueryRow(sqlStatement, username).Scan(&hash)
// 	return hash, err
// }

// func validatePass(db *sql.DB, username, senha string) bool {
// 	var (
// 		pass     []byte
// 		hashpass string
// 		err      error
// 	)

// 	pass = []byte(senha)
// 	hashpass, err = getHashPassword(db, username)
// 	if err != nil {
// 		log.Panic("Usuário não existente")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(hashpass), pass)
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }
