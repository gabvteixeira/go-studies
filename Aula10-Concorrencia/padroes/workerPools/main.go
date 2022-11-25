package main

import "fmt"

func main() {
	tarefas, resultados := make(chan int, 100), make(chan int, 100)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 1; i <= 100; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 1; i <= 100; i++ {
		fmt.Println(<-resultados)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}

	return fibonacci(n-2) + fibonacci(n-1)
}
