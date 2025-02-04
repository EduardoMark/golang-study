package main

import "fmt"

func init() {
	fmt.Println("Inicializing this function!")
}

func main() {
	fmt.Println("Eu não consigo ser o primeiro por causa do init")
}
