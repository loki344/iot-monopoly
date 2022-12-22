package boardDomain

import (
	"fmt"
	"iot-monopoly/communication"
)

type Player struct {
	Id        string `json:"id"`
	position  int
	AccountId string `json:"accountId"`
}

func NewPlayer(id string, position int, accountId string) *Player {
	return &Player{Id: id, position: position, AccountId: accountId}
}

func (p *Player) Position() int {
	return p.position
}

func (p *Player) SetPosition(position int) {

	if p.Position() == position {
		fmt.Println(fmt.Errorf("player already at position %d", position))
		return
	}

	fmt.Printf("MovePlayer player %s to fieldId %d\n", p.Id, position)
	p.position = position
	communication.FireEvent(communication.PLAYER_MOVED, NewPlayerMovedEvent(p.Id, position))
}
