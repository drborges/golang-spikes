package accounts

import "time"

type Account struct {
	Key      int
	Email    string
	Password string
}

type History struct {
	Key        int
	Action     string
	Actioner   Account
	ActionedAt time.Time
}

func GetById(accountId int) Account {
	return nil
}
