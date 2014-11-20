package main

import (
	"fmt"
	"time"
)

func producer(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func consumer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go producer(c)
	go consumer(c)

	var input string
	fmt.Scanln(&input)
}
