package domain

import (
	"fmt"
	"strconv"
)

type Account struct {
	id       string
	balance  int
	playerId string
}

func (a *Account) PlayerId() string {
	return a.playerId
}

func NewAccount(playerId string, index int) *Account {
	id := "Account_" + strconv.Itoa(index)
	return &Account{id: id, balance: 1_000, playerId: playerId}
}

func CreateUnlimitedAccount(playerId string) *Account {
	return &Account{id: "Bank", balance: 999_999, playerId: playerId}
}

func (a *Account) Id() string {
	return a.id
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Deposit(amount int) {
	fmt.Printf("Deposit %d to account %s\n", amount, a.id)
	a.balance += amount
}

func (a *Account) Pay(amount int) {
	fmt.Printf("Removing %d from account %s\n", amount, a.id)
	a.balance -= amount
}
