package models

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

type Usuario struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	FotoPerfil string `json:"fotoPerfil"`
	Username   string `json:"username"`
	UltimaVez  string `json:"ultimaVez"`
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
	Usuario{ID: 0, Nome: "Henriquezera", FotoPerfil: getImage(0)},
	Usuario{ID: 1, Nome: "Ericzera", FotoPerfil: getImage(1)},
	Usuario{ID: 2, Nome: "Dalton", FotoPerfil: getImage(2)},
	Usuario{ID: 3, Nome: "Dalton_2", FotoPerfil: getImage(2)},
}
