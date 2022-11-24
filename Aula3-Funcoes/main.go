package main

import (
	"errors"
	"fmt"
)

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func divisao(n1 int, n2 int) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("O denominador não pode ser zero")
	}
	return float64(n1 / n2), nil
}

func main() {
	fmt.Println(somar(1, 2))

	result, err := divisao(1, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
