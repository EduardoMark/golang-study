package main

import "fmt"

func recursion(num uint) uint {
	switch {
	case num == 0:
		return 1
	default:
		return recursion(num-1) * num
	}
}

func main() {
	result := recursion(5)
	fmt.Println(result)
}
