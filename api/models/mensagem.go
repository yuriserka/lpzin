package models

type Mensagem struct {
	ID        uint64 `json:"msgID"`
	ChatID    uint64 `json:"chatID"`
	Autor     int    `json:"autorID"`
	Conteudo  string `json:"conteudo"`
	HoraEnvio string `json:"hora_enviada"`
}
