package main

import "fmt"

func main() {
	// make(tipo, tamanho, capacidade)
	s := make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// só é possível atribuir dessa maneira até o indice limite
	s[9] = 10
	fmt.Println(s)
}