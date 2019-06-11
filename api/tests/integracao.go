package integracao

import (
	"fmt"
	"strconv"

	"github.com/yuriserka/lpzin/api/models"

	"github.com/yuriserka/lpzin/api/common"

	"github.com/yuriserka/lpzin/api/repositories"
)

const (
	entrar    = 1
	cadastrar = 2
	sair      = 3
)

var (
	repUsr = repositories.RepUser{}
	repCht = repositories.RepChat{}
)

func Init() {
	db, err := common.ConnDB()
	if err != nil {
		panic(fmt.Sprintf("db: %v", err))
	}

	defer db.Close()

	// schema.DropSchema(db)
	// schema.CreateSchema(db)

	repUsr.Init(db)
	repCht.Init(db)

	repCht.SetChat("Grupo de Testes", "Dev_Team")

	home()
}

func home() {
	var opt int
	for opt != 3 {
		fmt.Print("\t\tSistema de Vendas de Ingressos\n", "Escolha uma das opções abaixo:\n")
		fmt.Printf("[%d] Entrar\n[%d] Cadastrar\n[%d] Sair\n\topção: ", entrar, cadastrar, sair)

		switch fmt.Scanf("%d\n", &opt); opt {
		case entrar:
			if usr, logado := autenticar_test(); logado {
				enterChat_test(usr)
			}
		case cadastrar:
			cadastrar_test()
		}
	}
}

func enterChat_test(usr models.Usuario) {
	repCht.SetChatUser(strconv.Itoa(0), strconv.Itoa(usr.ID))
	for {
		fmt.Println("ta no chat", repCht.GetChat(strconv.Itoa(0)).Nome, "amigo")
		fmt.Println("Digite uma mensagem: ")
		fmt.Scanln(&msg)
	}
}

func autenticar_test() (models.Usuario, bool) {
	fmt.Println("\tAutenticar-se")
	fmt.Print("digite seu ID: ")
	var id string
	fmt.Scanln(&id)

	usuario := repUsr.GetUser(id)

	return usuario, true
}

func cadastrar_test() int {
	fmt.Println("\tCadastre-se")
	fmt.Print("Digite seu nome: ")
	var nome string
	fmt.Scanln(&nome)

	return repUsr.SetUser(nome, "FOTO DO USUARIO")
}
