package main

import "fmt"

// sem parâmetros e retorno
func first() {
	fmt.Println("My first function!")
}

// com parâmetros
func second(p1 int, p2 int) {
	fmt.Println( p1 + p2)
}

// com retorno
func third() bool {
	return false
}

// com parâmetros e retorno
func fourth(a,  b string) string {
	return a + b
}

// com mais de um retorno
func fifth() (string, int) {
	return "Hello", 17
}

func main() {
	first()

	second(17, 8)

	f3 := third()
	fmt.Println(f3)

	f4 := fourth("Quarta função", "Bom demais")
	fmt.Println(f4)

	r1, r2 := fifth()
	fmt.Println(r1, r2)
}