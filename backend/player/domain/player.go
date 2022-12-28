package playerDomain

import (
	"errors"
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

	//TODO move this code somewhere where it makes more sense
	totalFieldCount := 16
	if position > totalFieldCount-1 || position < 0 {
		errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", position, totalFieldCount))
	}

	if p.Position() == position {
		fmt.Println(fmt.Errorf("player already at position %d", position))
		return
	}

	if p.Position() > position && position < 5 {
		fmt.Println("Player completed a lap, creating lap finished")
		communication.FireEvent(communication.LAP_FINISHED, NewLapFinishedEvent(p.Id))
	}

	fmt.Printf("MovePlayer player %s to fieldId %d\n", p.Id, position)
	p.position = position
	communication.FireEvent(communication.PLAYER_MOVED, NewPlayerMovedEvent(p.Id, position))
}
