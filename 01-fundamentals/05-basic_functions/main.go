package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func main() {
	num1 := 17
	num2 := 8

	soma := somar(num1, num2)
	fmt.Println(soma)
}