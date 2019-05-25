package models

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

type Usuario struct {
	ID         int      `json:"id"`
	Nome       string   `json:"nome"`
	Msg        Mensagem `json:"mensagem"`
	FotoPerfil string   `json:"fotoPerfil"`
}

func getImage(id int) string {
	imgFile, err := os.Open("./static/img/foto_" + strconv.Itoa(id) + ".png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	size := fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return "data:image/png;base64, " + imgBase64Str
}

var Usuarios = []Usuario{
	Usuario{ID: 0, Nome: "Henriquezera", Msg: Mensagem{
		Autor:    0,
		Conteudo: "ola mundo_0",
	}, FotoPerfil: getImage(0)},
	Usuario{ID: 1, Nome: "Ericzera", Msg: Mensagem{
		Autor:    1,
		Conteudo: "ola mundo_1",
	}, FotoPerfil: getImage(1)},
	Usuario{ID: 2, Nome: "Dalton", Msg: Mensagem{
		Autor:    2,
		Conteudo: "ola mundo_2",
	}, FotoPerfil: getImage(2)},
	Usuario{ID: 3, Nome: "Dalton_2", Msg: Mensagem{
		Autor:    3,
		Conteudo: "ola mundo_2",
	}, FotoPerfil: getImage(2)},
}
