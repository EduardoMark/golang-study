package main

import "fmt"

func main() {
	// constantes
	const PI float64 = 3.14

	const (
		a = 1
		b = 2
	)

	const c, d = 3, 4

	// variáveis
	var hello string = "Hello"

	world := "World"

	var (
		e = true
		f = false
	)

	g, h := hello, world

	fmt.Println(hello, world, e, f, g, h)
}