package main

import "fmt"

func replace(p1, p2 int) (second int, first int) {
	first = p1
	second = p2

	return
}

func main() {
	a, b := replace(88, 44)
	fmt.Println(a, b)

	second, first := replace(55, 27)
	fmt.Println(second, first)
}