package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func main() {
	fmt.Println("struct")

	var u1 usuario
	u1.nome = "Darwin Nunez"
	u1.idade = 23
	fmt.Println(u1)

	u2 := usuario{"Valverde", 25}
	fmt.Println(u2)

	u3 := usuario{nome: "Luis Suárez"}
	fmt.Println(u3)
}
