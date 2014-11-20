package main

import "fmt"

func main() {
	// for ; i < 1000; { ... } => pre and post conditions may be ommited
	// for condition { ... }   => "for" is Go's "while"
	// for { ... }             => infinite loop
	for i := 0; i < 10; i++ {
		fmt.Println("Count:", i)
	}
}
