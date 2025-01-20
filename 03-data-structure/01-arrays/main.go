package main

import "fmt"

func main() {
    // único tipo de dados e tamanho estático
	notas := [3]float64{}

    fmt.Println(notas)

    // inserindo valores
    notas[0], notas[1], notas[2] = 10, 5.5, 8.7

    fmt.Println(notas)
    
}