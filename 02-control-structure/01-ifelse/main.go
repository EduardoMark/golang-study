package main

import "fmt"

func cnh(age int) {
	if age >= 18 {
		fmt.Println("Maior de idade")
	}

	if age >= 18 {
		fmt.Println("Permitido iniciar o processo de habilitação")
	} else {
		fmt.Println("Sem permição para iniciar o processo de habilitação!")
	}
}

func main(){
	cnh(21)
	cnh(18)
	cnh(16)
}