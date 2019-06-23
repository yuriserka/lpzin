package main

import (
    "fmt"
    "log"

    "github.com/yuriserka/lpzin/api/common"
    "github.com/yuriserka/lpzin/api/routers"
)

func main() {
    // testes.Init()

    db, err := common.ConnDB()
    if err != nil {
        log.Panic(fmt.Sprintf("db: %v", err))
    }

    defer db.Close()

    routers.Init(db)
    routers.Run()
}