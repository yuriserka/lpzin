package utils

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

// primeira função a ser executada no pacote
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ClearScreen limpa a tela do terminal.
func ClearScreen() {
	if executeClean, supportedPlatform := clear[runtime.GOOS]; supportedPlatform {
		executeClean()
	} else {
		panic("Plataforma não suportada, impossível limpar tela.")
	}
}
