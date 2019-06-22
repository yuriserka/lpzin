package repositorios

import (
	"database/sql"

	"github.com/yuriserka/lpzin/api/models"
)

func GetMensagem(msgID int) (*models.Mensagem, error) {
	sqlString := `
		select hr_env, msg, idchat, idusuario from mensagem
			where idmsg = $1 
	`
	var (
		horaEnviada, conteudo string
		chatid, autorid       int
	)
	row := db.QueryRow(sqlString, msgID)
	err := row.Scan(&horaEnviada, &conteudo, &chatid, &autorid)
	if err != nil {
		return nil, err
	}

	return &models.Mensagem{
		ID:        msgID,
		HoraEnvio: horaEnviada,
		Conteudo:  conteudo,
		ChatID:    chatid,
		Autor:     autorid,
	}, nil
}

func SetMensagem(conteudo string, chatID, autorID int) (int, error) {
	sqlString := `
		insert into mensagem (msg, idchat, idusuario) values ($1, $2, $3)
			returning idmsg
	`
	var id int
	err := db.QueryRow(sqlString, conteudo, chatID, autorID).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetMensagensUsuario(userID int) ([]*models.Mensagem, error) {
	sqlString := `
		select idmsg from mensagem where idusuario = $1 
	`
	rows, err := db.Query(sqlString, userID)
	if err != nil {
		return nil, err
	}

	mensagens, err := getMsgsFromRows(rows)
	if err != nil {
		return nil, err
	}

	return mensagens, nil
}

func getMsgsFromRows(rows *sql.Rows) ([]*models.Mensagem, error) {
	var msgs []*models.Mensagem
	for rows.Next() {
		var msgid int
		err := rows.Scan(&msgid)
		if err != nil {
			return nil, err
		}
		msg, err := GetMensagem(msgid)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}
