package main

import "fmt"

func main() {
	arr := [3]int{1 , 2, 3} // array
	sli := []int{4, 5, 6} // slice

	fmt.Println(arr, sli)

	// criando "fatias" (slices) de array
	fat := arr[0: 2]
	fmt.Println(fat)

	// adicionando elementos a uma slice
	sli = append(sli, 10, 11)
	fmt.Println(sli)
}