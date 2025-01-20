package main

import "fmt"

// só tem o for em golang

func main() {
	// como se fosse um while
	c := 1
	for c <= 10 {
		fmt.Println(c)
		c++
	}

	fmt.Println("---------------")

	// for tradicional
	for i := 0; i <= 20; i += 3 {
		fmt.Printf("%d\n", i)
	}

	// for infinito
	/*
		for {
			...
		}
	*/
}