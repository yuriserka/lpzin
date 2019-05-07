package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yuriserka/lpzin/routers"
)

// https://beego.me/docs/intro/ => olhar a documentação que explica bem
// é bom dar um go get ./... antes de mexer só por garantia
// pra rodar só fazer "bee run"

func main() {
	beego.Run()
}
