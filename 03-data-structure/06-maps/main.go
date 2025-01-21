package main

import "fmt"

func main() {
	// approved := make(map[int]string)
	approved := map[int]string{}

	approved[123] = "John"
	approved[321] = "Doe"
	approved[988] = "Ane"

	fmt.Println(approved)

	for id, name := range approved {
		fmt.Printf("%s (ID: %d)\n", name, id)
	}

	delete(approved, 988)
}