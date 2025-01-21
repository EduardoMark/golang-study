package main

import "fmt"

func main() {
	employeesByLetter := map[string]map[string]float64{
		"A": {
			"Ana": 1234.56,
		},
		"B": {
			"Bruno": 3254.65,
		},
	}

	fmt.Println(employeesByLetter)

	employeesByLetter["C"] = map[string]float64{
		"Carlos": 1234.50,
	} 

	fmt.Println(employeesByLetter)
}