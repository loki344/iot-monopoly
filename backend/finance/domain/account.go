package financeDomain

import "fmt"

type Account struct {
	Id       string `json:"id"`
	PlayerId string `json:"playerId"`
	Balance  int    `json:"balance"`
}

func (a *Account) Add(amount int) {
	fmt.Printf("Add %d to account %s\n", amount, a.Id)

	a.Balance += amount
}

func (a *Account) Subtract(amount int) {
	fmt.Printf("Removing %d from account %s\n", amount, a.Id)
	a.Balance -= amount
}
