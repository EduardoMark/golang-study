package main

import (
	"fmt"
	"time"
)

// se tentar ler um dado depois que a goroutine acabar gera um deadlock
func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 17 // ação bloqueante
	fmt.Println("Só executo depois do channel ser lido!")
}

func main() {
	c := make(chan int)

	go routine(c)    // criando uma goroutine
	fmt.Println(<-c) // lendo o channel

	fmt.Println("Hello World")

	fmt.Println(<-c) // error -> deadlock
}
