package schema

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	sqlStatement string = `
	CREATE TABLE IF NOT EXISTS Chat (
		ID				SERIAL NOT NULL PRIMARY KEY,
		Nome			VARCHAR(100) NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS Usuario (
		ID				SERIAL NOT NULL PRIMARY KEY,
		Nome			VARCHAR(100) NOT NULL,
		Foto			BYTEA NOT NULL,
		username		VARCHAR(100) UNIQUE,
		senha			VARCHAR(4000),
		ultima_vez		TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC+3')
	);
	
	CREATE TABLE IF NOT EXISTS Chat_tem_Usuario (
		IDChat			INTEGER NOT NULL REFERENCES Chat(ID),
		IDUsuario		INTEGER NOT NULL REFERENCES Usuario(ID),
		PRIMARY KEY(IDChat, IDUsuario)
	);
	
	CREATE TABLE IF NOT EXISTS Mensagem (
		IDMsg			SERIAL NOT NULL PRIMARY KEY,
		Hr_env			TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC+3') NOT NULL,
		Msg				VARCHAR(500) NOT NULL,
		IDChat			SERIAL NOT NULL REFERENCES Chat(ID),
		IDUsuario		SERIAL NOT NULL REFERENCES Usuario(ID)
	);
	`
	sqlDrop string = `
	DROP TABLE IF EXISTS Chat CASCADE;
	DROP TABLE IF EXISTS Usuario CASCADE;
	DROP TABLE IF EXISTS Chat_tem_usuario CASCADE;
	DROP TABLE IF EXISTS Mensagem CASCADE;`
)

// CreateSchema executa o sql para criar as tabelas no banco de dados
func CreateSchema(db *sql.DB) {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Printf("Error %+v\n", err)
	}
	fmt.Println("Schema criado com sucesso")
}

// DropSchema é uma função temporária para realizar testes
func DropSchema(db *sql.DB) {
	_, err := db.Exec(sqlDrop)
	if err != nil {
		log.Printf("Error %+v\n", err)
	}
	fmt.Println("Schema dropado com sucesso")
}
