package models

type Mensagem struct {
	Autor    int    `json:"autorID"`
	Conteudo string `json:"conteudo"`
}
