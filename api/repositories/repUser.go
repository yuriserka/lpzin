package repositories

import (
	"database/sql"
	"log"
)

type RepUser struct {
	db *sql.DB
}

func (rep *RepUser) Init(db *sql.DB) {
	rep.db = db
}

func (rep *RepUser) SetUser(nome string, foto string) {
	sqlStatement := `
	INSERT INTO Usuario (nome, foto)
	VALUES ($1, $2);`

	_, err := rep.db.Exec(sqlStatement, nome, foto)
	if err != nil {
		log.Printf("Error %+v\n", err)
	}
}

// func (rep *RepUser) GetUser(userid string) (models.Usuario, error) {
// 	sqlStatement := `SELECT id, `
// }
