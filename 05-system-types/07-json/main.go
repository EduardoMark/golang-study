package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := People{"Eduardo", 21}

	// struct to json
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json to struct
	var p2 People
	jsonString := `{"name":"Dudu", "age":21}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Age, p2.Name)
}
