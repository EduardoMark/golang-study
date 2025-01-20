package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// inicia variável acessivel somente no bloco if e seus decendentes
	if i := randomNumber(); i > 5 {
		fmt.Println("You Win!!!!")
	} else {
		fmt.Println("Oh You Lost =(")
	}
}