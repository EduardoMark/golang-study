package main

import "fmt"

func main() {
	// não é possível usar o "append" em array

	slice1 := []int{}
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(slice1)

	// só é possível copiar uma slice ou string
	s2 := make([]int, 2)
	copy(s2, slice1) // ambas slices devem ser do mesmo tipo

	fmt.Println(s2)
}