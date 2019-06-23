package main

import (
	"github.com/yuriserka/lpzin/api/testes"
)

func main() {
	testes.Init()

	// db, err := common.ConnDB()
	// if err != nil {
	// 	log.Panic(fmt.Sprintf("db: %v", err))
	// }

	// defer db.Close()

	// routers.Init(db)
	// routers.Run()
}
