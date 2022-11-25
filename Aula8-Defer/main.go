package main

import (
	"fmt"
	"time"
)

func funcao1() {
	defer fmt.Println("fim da execucao")
	fmt.Println("Executando funcao 1")
	time.Sleep(5)
	fmt.Println("..............")
}

func funcao2() {
	fmt.Println("Executando funcao 2")
}

func main() {
	defer funcao1()
	funcao2()
}
