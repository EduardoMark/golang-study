package main

import "fmt"

func obterValorAprovado(num int) int {
	defer fmt.Println("End...")
	if num > 100 {
		fmt.Println("Maior que 100")
		return num
	}

	fmt.Println("Menor ou igual a 100")
	return num
}

func main() {
	fmt.Println(obterValorAprovado(100))
}
