package main

import "fmt"

func closure() func() {
	name := "Eduardo"
	x := func() {
		fmt.Println("Eu estou acessando o local onde fui declarada: ", name)
	}
	return x
}

func main() {
	name := "Marques"
	fmt.Println(name)
	printName := closure()
	printName()
}
