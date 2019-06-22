package repositorios

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yuriserka/lpzin/schema"

	"github.com/yuriserka/lpzin/api/common"
)

var db *sql.DB

func init() {
	var e error
	db, e = common.ConnDB()
	if e != nil {
		log.Panic(fmt.Sprintf("Error %+v\n", e))
	}
}

func DropaEsquema() {
	schema.DropSchema(db)
}

func CriaEsquema() {
	schema.CreateSchema(db)
}
