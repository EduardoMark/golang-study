package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

// essa declaração de variável := não é possivel fora de um bloco de código

// sub := func(a, b int) int {
// 	return a - b
// } 

func main() {
	fmt.Println(sum(17, 8))

	sub := func(a, b int) int {
		return a - b
	} 

	fmt.Println(sub(8, 17))
}