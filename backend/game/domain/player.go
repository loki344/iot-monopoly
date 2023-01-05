package domain

import (
	"strconv"
)

type Player struct {
	id                        string
	position                  int
	account                   *Account
	escapeFromPrisonCardCount int
}

func (p *Player) EscapeFromPrisonCardCount() int {
	return p.escapeFromPrisonCardCount
}

func (p *Player) Account() *Account {
	return p.account
}

func (p *Player) Id() string {
	return p.id
}

func NewPlayer(index int) *Player {
	id := "Player_" + strconv.Itoa(index)
	return &Player{id: id, position: 1, account: NewAccount(id, index)}
}

func CreateBank() *Player {
	return &Player{id: "Bank", position: 0, account: CreateUnlimitedAccount("Bank")}
}

func (p *Player) Position() int {
	return p.position
}
