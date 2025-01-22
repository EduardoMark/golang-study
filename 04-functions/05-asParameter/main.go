package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func exec(function func(int, int) int, a, b int) int {
	return function(a, b)
}

func main() {
	ex1 := exec(multiply, 70, 7)
	ex2 := exec(division, 100, 10)

	fmt.Println(ex1, ex2)
}