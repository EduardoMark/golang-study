package main

import "fmt"

type vehicle struct {
	name        string
	maxVelocity int
}

type car struct {
	vehicle // adiciona campos anônimos
	carType string
	price   float64
}

func main() {
	c := car{}
	c.name = "Cronos"
	c.maxVelocity = 220
	c.carType = "Sedan"
	c.price = 60000.00

	fmt.Println(c)
}
