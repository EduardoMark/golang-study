package main

import (
	"fmt"
	"math"
)

func main() {
	a := 100
	b := 50

	fmt.Println("Soma -> ", a + b)
	fmt.Println("Subtração -> ", a - b)
	fmt.Println("Multiplicação -> ", a * b)
	fmt.Println("Divisão -> ", a / b)
	fmt.Println("Módulo -> ", a % b)

	// bitwise
	fmt.Println("E ->", a & b) // 11 & 10 = 10
	fmt.Println("Ou ->", a & b) // 11 | 10 = 11
	fmt.Println("Xor ->", a & b) // 11 ^ 10 = 01

	fmt.Println(math.Pow(2.0, 2.0)) // potência
}