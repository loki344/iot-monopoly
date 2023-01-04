package domain

import (
	"fmt"
	"strconv"
)

type Player struct {
	id       string
	position int
	account  *Account
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

func (p *Player) SetPosition(position int) {

	if p.Position() == position {
		fmt.Println(fmt.Errorf("player already at position %d", position))
		return
	}
	p.position = position
}
