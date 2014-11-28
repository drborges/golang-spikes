package main

import "fmt"

func add(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	a := 1
	b := 2
	result := add(a, b)
	fmt.Printf("%v + %v = %v", a, b, result)
}
