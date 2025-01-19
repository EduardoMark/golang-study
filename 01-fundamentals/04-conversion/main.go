package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// fmt.Println(x / y) // isso não é possível
	fmt.Println(x / float64(y))

	nota := 7.5
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(85)) // retorna o valor da tabela Unicode

	// string para int
	num := strconv.Itoa(22)
	fmt.Println("O num é " + num)
}