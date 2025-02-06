package main

import "fmt"

// interfaces fornece um conjunto de métodos que se espera que uma struct tenha implementado
type Animal interface {
	Onomatopeia() string
}

type Cachorro struct {
	name string
}

func (c Cachorro) Onomatopeia() string {
	return "Auau"
}

func run(animal Animal) {
	fmt.Println(animal.Onomatopeia())
}

func main() {
	c := Cachorro{"Thor"}
	fmt.Println(c.Onomatopeia())
}
