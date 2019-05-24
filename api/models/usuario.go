package models

type Usuario struct {
	ID   int    `json:"id"`
	Nick string `json:"username"`
	Msg  string `json:"message"`
}

var Usuarios = []Usuario{
	Usuario{ID: 0, Nick: "Usr_0", Msg: "ola mundo0"},
	Usuario{ID: 1, Nick: "Usr_1", Msg: "ola mundo1"},
	Usuario{ID: 2, Nick: "Usr_2", Msg: "ola mundo2"},
}
