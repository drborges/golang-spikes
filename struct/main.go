package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Say(message string) string {
	return message + ", I'm " + p.Name + ", and I am " + strconv.Itoa(p.Age) + " years old"
}

func main() {
	var p = Person{Name: "Diego", Age: 29}
	fmt.Printf(p.Say("Hello there..."))
}
