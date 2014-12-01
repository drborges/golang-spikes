package main

import (
	"fmt"
	"models/accounts"
)

func main() {
	account := accounts.GetById(123)
	fmt.Println(account.Email)
}
