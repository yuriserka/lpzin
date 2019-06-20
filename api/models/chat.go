package models

type Chat struct {
	ID        int        `json:"chatID"`
	Nome      string     `json:"chatName"`
	Users     []Usuario  `json:"users"`
	Mensagens []Mensagem `json:"messages"`
}

var Chats = []Chat{
	Chat{ID: 0, Nome: "Klub dos WebAmigos", Users: Usuarios, Mensagens: []Mensagem{}},
	Chat{ID: 2, Nome: "XD", Users: Usuarios, Mensagens: []Mensagem{
		Mensagem{ID: 5, ChatID: 2, Autor: 1, Conteudo: "kkj"},
	}},
}
