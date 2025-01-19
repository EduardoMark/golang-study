package main

import "fmt"

func main() {
	var a byte = 10
	fmt.Println(a)

	i := 17 // inferência de tipo
	i += 2
	i -= 2
	i /= 2
	i *= 2
	i %= 1
	i++
	i--

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}