package main

import "fmt"

func main() {
	ages := [...]int{12, 44, 52, 21, 18}

	for i, age := range ages {
		fmt.Printf("%d) %d\n", i, age)
	}

	fmt.Println("--------")

	for _, age := range ages {
		fmt.Println(age)
	}
}