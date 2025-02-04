package main

import (
	"fmt"
	"strings"
)

type people struct {
	Name    string
	Surname string
}

func (p people) getFullName() string {
	return p.Name + " " + p.Surname
}

// passando como ponteiro para alterar o valor da variável externa
func (p *people) setFullName(fullname string) {
	slice := strings.Split(fullname, " ")
	p.Name = slice[0]
	p.Surname = slice[1]
}

func main() {
	p1 := people{
		Name:    "Eduardo",
		Surname: "Marques",
	}

	fmt.Println(p1.getFullName())

	p1.setFullName("John Doe")
	fmt.Println(p1.getFullName())
}
