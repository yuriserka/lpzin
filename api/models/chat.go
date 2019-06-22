package models

type Chat struct {
	ID               int         `json:"ID"`
	Nome             string      `json:"Nome"`
	Users            []*Usuario  `json:"Usuarios"`
	Mensagens        []*Mensagem `json:"Mensagens"`
	UsuarioCriadorID int         `json:"UsuarioCriadorID"`
}
