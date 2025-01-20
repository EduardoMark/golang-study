package main

import (
	"fmt"
	"time"
)

func noteForConcept(n float64) string {
	note := int(n)

	switch note {
	case 10:
		fallthrough
	case 9: 
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida!"	
	}
}

func greeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}

func main() {
	fmt.Println(noteForConcept(10))
	fmt.Println(noteForConcept(7))
	fmt.Println(noteForConcept(3))

	greeting()
}