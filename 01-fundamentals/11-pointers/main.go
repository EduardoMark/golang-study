package main

import "fmt"

func main() {
	c := 8

	var p *int = nil

	p = &c // pegando o endereço da memória da variável c
	
	fmt.Println(p) // endereço da memória
	fmt.Println(*p) // valor que está armazenado
	fmt.Println(c) // valor da variável
}