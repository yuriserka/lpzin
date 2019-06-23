package utils

import (
	"fmt"
)

// Pause simula a pausa da aplicação até que a tecla "ENTER" seja digitada.
func Pause() {
	fmt.Println("Digite Enter para continuar...")
	fmt.Scanln()
}
