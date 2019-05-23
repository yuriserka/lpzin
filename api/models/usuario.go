package models

type Usuario struct {
	ID   int    `json:"id"`
	Nick string `json:"username"`
	Msg  string `json:"message"`
}

var (
	Db       map[int]Usuario
	Contador int
)

func init() {
	Db = make(map[int]Usuario)
	Contador = 0
}
