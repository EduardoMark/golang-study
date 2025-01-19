package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// sem sinal (positivos) uint8, uint 16, uint32, uint64
	var a byte = 255
	fmt.Println(reflect.TypeOf(a))

	// com sinal int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é:", i1)

	// tipo rune, mapeia a tabela Unicode
	var i2 rune = 'a'
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2)

	// inteiros: -100, -1, 10, 1000
	// floats: -1.45, 8.17
	// bool: true, false
	// string: "Hello World!" 
}