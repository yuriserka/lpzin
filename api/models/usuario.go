package models

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

type Usuario struct {
	ID         int    `json:"ID"`
	Nome       string `json:"Nome"`
	FotoPerfil string `json:"FotoPerfil"`
	Username   string `json:"Username"`
	UltimaVez  string `json:"UltimaVez"`
	Senha      string `json:"Senha"`
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
