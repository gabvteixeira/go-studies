package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Teixeira",
	}

	fmt.Println(usuario["nome"])
}
