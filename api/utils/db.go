package utils

//Fazer o schema do DB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Informações para conectar no db postgree
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "chat_lp"
)

//CreateConnString Prepara as informações para passar para a função que conecta o banco de dados
func CreateConnString(host string, port int, user string, password string, dbname string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

//ConnDB conecta com o bd e testa a conexão com o bd
func ConnDB() (*sql.DB, error) {
	connectionString := CreateConnString(host, port, user, password, dbname)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("db: %v", err))
	}

	err = conn.Ping()
	if err != nil {
		panic(fmt.Sprintf("db: %v", err))
	}

	fmt.Println("Conexão criada com sucesso")

	return conn, err
}

// usar apenas uma conexão pro DB, para realizar as operações
// usando o defer db.Close() para fechar o DB ao fim do escopo da função
// func DesconnDB(db *sql.DB) error {
// 	err := db.Close()

// 	if err != nil {
// 		panic(fmt.Sprintf("db: %v", err))
// 	}
// 	return err
// }
