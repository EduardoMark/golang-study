package main

import "fmt"

func dayOfWeek(day int) string {
	if day == 1 {
		return "Domingo"
	} else if day == 2 {
		return "Segunda"
	} else if day == 3 {
		return "Terça"
	} else if day == 4 {
		return "Quarta"
	} else if day == 5 {
		return "Quinta"
	} else if day == 6 {
		return "Sexta"
	} else if day == 7 {
		return "Sabádo"
	} else {
		return "Dia inválido!"
	}
} 

func main() {
	day := dayOfWeek(6)
	fmt.Println("Dia:", day)
}