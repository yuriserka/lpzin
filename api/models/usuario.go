package models

type Usuario struct {
	ID         int    `json:"ID"`
	Nome       string `json:"Nome"`
	FotoPerfil string `json:"FotoPerfil"`
	Username   string `json:"Username"`
	UltimaVez  string `json:"UltimaVez"`
	Senha      string `json:"Senha"`
}
