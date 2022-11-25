package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(<-generator("Patio Solar"))
	}
}

func generator(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- texto
			time.Sleep(time.Second)
		}
	}()

	return canal
}
