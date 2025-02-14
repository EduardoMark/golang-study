package main

import (
	"fmt"
	"time"
)

// para criar uma goroutine adiciona "go" antes da func/method

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go numbers()
	time.Sleep(time.Second * 5)
	fmt.Println("\nFim da execução")
}
