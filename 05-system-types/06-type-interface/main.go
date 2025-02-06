package main

import (
	"fmt"
	"reflect"
)

func print(x interface{}) {
	fmt.Println(reflect.TypeOf((x)))
}

func main() {
	print(10)
	print("EU")
	print(true)
}
