package main

import (
	"fmt"
	"time"
)

// channel é a forma de comunicação entre goroutines
func multiply(num int, c chan int) {
	time.Sleep(time.Second)
	c <- num * 2 // canal está recebendo dado

	time.Sleep(time.Second)
	c <- num * 3

	time.Sleep(time.Second)
	c <- num * 4
}

func main() {
	c := make(chan int) // criando um a channel

	go multiply(10, c) // criando uma goroutine

	x, y := <-c, <-c // recebendo dados do channel
	fmt.Println(x, y)
	fmt.Println(<-c)

	fmt.Println("Fim da execução!")
}
