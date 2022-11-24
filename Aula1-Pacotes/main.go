package main

import (
	"Aula1-Pacotes/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	fmt.Println("Validando email...")
	err := checkmail.ValidateFormat("devbookgmail.com")

	if err != nil {
		fmt.Println("E-mail inválido")
	}
}
