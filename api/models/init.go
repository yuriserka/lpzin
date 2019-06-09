package models

var (
	Db       map[int]Usuario
	Contador int
)

func init() {
	Db = make(map[int]Usuario)
	Contador = 0
}
