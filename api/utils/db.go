package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Informações para conectar no db postgree
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "33583549"
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
