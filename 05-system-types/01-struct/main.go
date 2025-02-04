package main

import "fmt"

// struct {identificador - tipo}
type people struct {
	name string
	age  int
	job  string
}

// Método: função com receiver (receptor)
func (p people) peopleInfo() string {
	return fmt.Sprintf("%s é %s, e tem %d anos", p.name, p.job, p.age)
}

func main() {
	people1 := people{
		name: "Eduardo",
		age:  21,
		job:  "Desenvolvedor",
	}

	fmt.Println(people1)
	fmt.Println(people1.peopleInfo())
}
