package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Média: %.2f\n", average(5.6, 7.8, 9.6, 6, 7))
}
