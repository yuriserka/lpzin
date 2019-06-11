package integracao

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yuriserka/lpzin/api/models"
	"github.com/yuriserka/lpzin/schema"

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
	repMsg = repositories.RepMessage{}
)

// Init simula de forma simplificada a interação com o banco de dados do sistema a fim de testar suas
// funcionalidades
func Init() {
	db, err := common.ConnDB()
	if err != nil {
		panic(fmt.Sprintf("db: %v", err))
	}

	defer db.Close()

	schema.DropSchema(db)
	schema.CreateSchema(db)

	repUsr.Init(db)
	repCht.Init(db)
	repMsg.Init(db)

	idAdm := repUsr.SetUser("AdmCorno", "kkk bovino")
	repCht.SetChat("Grupo de Testes", strconv.Itoa(idAdm))

	home()
}

func home() {
	var opt int
	for opt != 3 {
		fmt.Print("\t\tTeleZap simples\n", "Escolha uma das opções abaixo:\n")
		fmt.Printf("[%d] Entrar\n[%d] Cadastrar\n[%d] Sair\n\topção: ", entrar, cadastrar, sair)

		switch fmt.Scanf("%d\n", &opt); opt {
		case entrar:
			if usr, logado := autenticarTest(); logado {
				entrarNoChatTest(usr)
			}
		case cadastrar:
			cadastrarTest()
		}
	}
}

func entrarNoChatTest(usr models.Usuario) {
	const chatIDStub = 1
	repCht.SetChatUser(strconv.Itoa(chatIDStub), strconv.Itoa(usr.ID))
	chat := repCht.GetChat(strconv.Itoa(chatIDStub))
	for {
		msgs := repCht.GetChatMsgs(strconv.Itoa(chatIDStub))

		fmt.Printf("\tta no chat %4s amigo\n", chat.Nome)

		for _, v := range msgs {
			u := repUsr.GetUser(strconv.Itoa(v.Autor))
			fmt.Printf("%s => %-6s %-10s\n", u.Nome, v.Conteudo, v.HoraEnvio)
		}

		fmt.Print("Digite uma mensagem: ")
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		msg = msg[:len(msg)-2] // retira o '\n' da mensagem

		repMsg.SetMsg(strconv.Itoa(usr.ID), strconv.Itoa(chat.ID), msg)
	}
}

func autenticarTest() (models.Usuario, bool) {
	fmt.Println("\tAutenticar-se")
	fmt.Print("digite seu ID: ")
	var id string
	fmt.Scanln(&id)

	usuario := repUsr.GetUser(id)

	return usuario, true
}

func cadastrarTest() int {
	fmt.Println("\tCadastre-se")
	fmt.Print("Digite seu nome: ")
	var nome string
	fmt.Scanln(&nome)

	return repUsr.SetUser(nome, "FOTO DO USUARIO")
}
