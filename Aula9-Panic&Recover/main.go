package main

import "fmt"

func recuperarFuncao1() {
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado")
	}
}

func funcao1() {
	defer recuperarFuncao1()
	panic("Programa entrando em panico!!!")
}

func main() {
	funcao1()
}
