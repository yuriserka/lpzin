package models

type Mensagem struct {
	ID        int    `json:"ID"`
	ChatID    int    `json:"ChatID"`
	Autor     int    `json:"AutorID"`
	Conteudo  string `json:"Conteudo"`
	HoraEnvio string `json:"HoraEnviada"`
}
