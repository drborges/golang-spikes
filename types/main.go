package main

import "fmt"

type Log string

func (l Log) Info(msg string) string {
	return fmt.Sprint("[", l, "] ", msg)
}

func main() {
	var log = Log("MyLogTag").Info("Well hello there...")

	fmt.Println(log)
}
