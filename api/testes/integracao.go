package testes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yuriserka/lpzin/api/common"
	"github.com/yuriserka/lpzin/api/repositories"
	"github.com/yuriserka/lpzin/schema"

	"github.com/yuriserka/lpzin/api/utils"
)

var (
	db, err = common.ConnDB()
	repChat = &repositories.RepChat{}
	repUser = &repositories.RepUser{}
	repMsg  = &repositories.RepMessage{}
)

// Init simula de forma simplificada a interação com o banco de dados do sistema a fim de testar suas
// funcionalidades
func Init() {
	if err != nil {
		log.Panic(fmt.Sprintf("db: %v", err))
	}

	schema.DropSchema(db) // método temporário para realizar testes
	schema.CreateSchema(db)

	defer db.Close()
	repChat.Init(db)
	repUser.Init(db)
	repMsg.Init(db)
	home()
}

func home() {

	const (
		klogar     = iota + 1
		kcadastrar = iota + 1
		ksair      = iota + 1
	)
	menuLogin := map[int]string{
		klogar:     "Logar",
		kcadastrar: "Cadastrar-se",
		ksair:      "Sair",
	}

	opt := -1
	sortedIndexesLogin := utils.OrdenaMap(menuLogin)
	for opt != ksair {
		utils.ClearScreen()
		fmt.Println("\tTeste de Integração")
		for _, k := range sortedIndexesLogin {
			fmt.Printf("[%d] %s\n", k, menuLogin[k])
		}
		fmt.Print("\tOpção: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case klogar:
			if logado, id := logarNoWebZap(); logado {
				logadoHome(id)
			}
		case kcadastrar:
			cadastrarTest()
		}
	}
}

func logadoHome(id int) {
	const (
		kusar                 = iota + 1
		kentrar               = iota + 1
		kcriarChat            = iota + 1
		kgetMsgs              = iota + 1
		kgetChatsParticipante = iota + 1
		kvoltar               = iota + 1
	)
	menu := map[int]string{
		kusar:                 "Usar o Chat",
		kentrar:               "Entrar em um Chat",
		kcriarChat:            "Criar Chat",
		kgetMsgs:              "Ver Mensagens",
		kgetChatsParticipante: "Ver Chats que participa",
		kvoltar:               "Voltar",
	}
	sortedIndexes := utils.OrdenaMap(menu)
	opt := -1
	for opt != kvoltar {
		utils.ClearScreen()
		fmt.Println("\tTeste dos Serviços")
		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}
		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case kentrar:
			entrarNoChatTest(id)
		case kusar:
			usarChatTest(id)
		case kcriarChat:
			criarChatTest(id)
		case kgetMsgs:
			getUserMsgsTest(id)
		case kgetChatsParticipante:
			getUserChatsIDTest(id)
		}
	}
}

func usarChatTest(id int) {
	var chatIDStub int
	fmt.Printf("Digite o id do chat: ")
	fmt.Scanf("%d\n", &chatIDStub)

	chat, err := repChat.GetChat(chatIDStub)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		msgs, err := repChat.GetChatMsgs(chatIDStub)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		utils.ClearScreen()
		fmt.Printf("\tta no chat %4s amigo\n", chat.Nome)
		for _, v := range msgs {
			u, _ := repUser.GetUser(v.Autor)
			fmt.Printf("%s => %-6s %-10s\n", u.Nome, v.Conteudo, v.HoraEnvio)
		}

		fmt.Print("Digite uma mensagem: ")
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		msg = msg[:len(msg)-2] // retira o '\n' da mensagem

		if c := strings.Compare(msg, "quit"); c == 0 {
			break
		}

		_, errr := repMsg.SetMsg(id, chatIDStub, msg)
		if errr != nil {
			fmt.Println(errr.Error())
			return
		}
	}
}

func entrarNoChatTest(id int) {
	var chatIDStub int
	fmt.Printf("Digite o id do chat que deseja entrar: ")
	fmt.Scanf("%d\n", &chatIDStub)

	if err := repChat.SetUserInChat(chatIDStub, id); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func cadastrarTest() int {
	fmt.Println("\tCadastre-se")
	fmt.Print("Digite seu nome: ")
	var (
		nome     string
		username string
		senha    string
		reader   *bufio.Reader
	)
	reader = bufio.NewReader(os.Stdin)
	nome, _ = reader.ReadString('\n')
	nome = nome[:len(nome)-2]

	fmt.Print("Digite seu username: ")
	reader = bufio.NewReader(os.Stdin)
	username, _ = reader.ReadString('\n')
	username = username[:len(username)-2]

	fmt.Print("Digite sua senha: ")
	reader = bufio.NewReader(os.Stdin)
	senha, _ = reader.ReadString('\n')
	senha = senha[:len(senha)-2]

	id, e := repUser.SetUser(nome, "", username, senha)
	if e != nil {
		fmt.Println(e.Error())
	}
	return id
}

func getUserMsgsTest(id int) {
	if id < 0 {
		fmt.Println("cadastre-se primeiro")
		return
	}
	msgs, err := repUser.GetUserMsgs(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range msgs {
		fmt.Printf("%d => %v %v\n", v.Autor, v.Conteudo, v.HoraEnvio)
	}
}

func getUserChatsIDTest(id int) {
	if id < 0 {
		fmt.Println("cadastre-se primeiro")
		return
	}
	chats, err := repUser.GetUserChats(id)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}
	for _, v := range chats {
		fmt.Printf("ID: %d => %v\n", v.ID, v.Nome)
	}
	return
}

func criarChatTest(id int) {
	var (
		reader *bufio.Reader
		nome   string
		// userid int
	)
	fmt.Print("Digite o nome do chat: ")
	reader = bufio.NewReader(os.Stdin)
	nome, _ = reader.ReadString('\n')
	nome = nome[:len(nome)-2]

	// fmt.Println()
	// fmt.Print("Digite o id do criador do chat: ")
	// reader = bufio.NewReader(os.Stdin)
	// id, _ := reader.ReadString('\n')
	// userid, _ = strconv.Atoi(id[:len(id)-2])

	chid, err := repChat.SetChat(nome)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = repChat.SetUserInChat(chid, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func logarNoWebZap() (bool, int) {
	var (
		username string
		password string
		id       int
	)
	opt := "n"
	for opt == "n" {
		fmt.Print("Digite seu username: ")
		reader := bufio.NewReader(os.Stdin)
		username, _ = reader.ReadString('\n')
		username = username[:len(username)-2]

		fmt.Print("Digite sua senha: ")
		reader = bufio.NewReader(os.Stdin)
		password, _ = reader.ReadString('\n')
		password = password[:len(password)-2]

		logado, err := repUser.UserAuth(username, password)
		if err != nil {
			if !logado {
				fmt.Printf("Username ou senha errada\n")
				fmt.Print("Quer sair (s/n): ")
				reader = bufio.NewReader(os.Stdin)
				opt, _ = reader.ReadString('\n')
				opt = opt[:len(opt)-2]
			}
		}
		if logado {
			id, _ = repUser.GetUserID(username)
			return logado, id
		}
	}
	return false, -1
}
