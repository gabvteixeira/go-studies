package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{}

	fmt.Println(array2)

	slice := []int{1}

	fmt.Println(slice)
}
