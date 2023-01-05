package domain

import (
	"strconv"
)

type Player struct {
	id              string
	position        int
	account         *Account
	prisonCardCount int
}

func (p *Player) EscapeFromPrisonCardCount() int {
	return p.prisonCardCount
}

func (p *Player) Account() *Account {
	return p.account
}

func (p *Player) Id() string {
	return p.id
}

func (p *Player) Position() int {
	return p.position
}

func (p *Player) IncreasePrisonCardCount() {
	p.prisonCardCount = p.prisonCardCount + 1
}
func (p *Player) DecreasePrisonCardCount() {
	p.prisonCardCount = p.prisonCardCount - 1
}

func newPlayer(index int) *Player {
	id := "Player_" + strconv.Itoa(index)
	return &Player{id: id, position: 1, account: newAccount(id, index)}
}
