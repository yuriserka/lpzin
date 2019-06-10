package models

type Mensagem struct {
	ID        int    `json:"msgID"`
	ChatID    int    `json:"chatID"`
	Autor     int    `json:"autorID"`
	Conteudo  string `json:"conteudo"`
	HoraEnvio string `json:"hora_enviada"`
}
