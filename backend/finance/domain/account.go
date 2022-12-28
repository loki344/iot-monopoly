package financeDomain

import "fmt"

type Account struct {
	id       string
	playerId string
	balance  int
}

func NewAccount(id string, playerId string, balance int) *Account {
	return &Account{id: id, playerId: playerId, balance: balance}
}

func (a *Account) Id() string {
	return a.id
}

func (a *Account) PlayerId() string {
	return a.playerId
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Add(amount int) {
	fmt.Printf("Add %d to account %s\n", amount, a.id)

	a.balance += amount
}

func (a *Account) Subtract(amount int) {
	fmt.Printf("Removing %d from account %s\n", amount, a.id)
	a.balance -= amount
}
