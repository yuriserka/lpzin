package models

type Chat struct {
	ID        uint64     `json:"chatID"`
	Nome      string     `json:"chatName"`
	Users     []Usuario  `json:users`
	Mensagens []Mensagem `json:messages`
}

var Chats = []Chat{
	Chat{ID: 0, Nome: "Klub dos WebAmigos", Users: Usuarios, Mensagens: []Mensagem{
		Mensagem{ID: 0, ChatID: 0, Autor: 0, Conteudo: "Oiii"},
		Mensagem{ID: 1, ChatID: 0, Autor: 1, Conteudo: "kk eae men"},
	}},
	Chat{ID: 2, Nome: "XD", Users: Usuarios, Mensagens: []Mensagem{
		Mensagem{ID: 5, ChatID: 2, Autor: 1, Conteudo: "kkj"},
	}},
}
