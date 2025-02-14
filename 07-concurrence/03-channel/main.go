package main

import "fmt"

// channel é um tipo em go
// chan

func main() {
	ch := make(chan int, 1)

	// enviando dados para o canal (escrita)
	ch <- 1

	// canal enviando dados do canal (leitura)
	r := <-ch
	fmt.Println(r)
}
