package boardDomain

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
)

type Player struct {
	Id int
	// TODO maybe change position to field
	Position int
	Balance  int
}

func (player *Player) MovePlayer(fieldIndex int) {

	//TODO get rid of magic numbers 10!!
	fmt.Printf("Move player %d to fieldIndex %d\n", player.Id, fieldIndex)
	if (player.Position >= 10 && player.Position < 16) && (fieldIndex >= 0 && fieldIndex <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		ts := bus.GetBus()
		handler, err := ts.RequestOnce("lapFinished", LapFinishedEvent{player.Id})
		if err != nil {
			//TODO something went wrong
			return
		}
		handler.Fire()
	}

	player.Position = fieldIndex
}

type LapFinishedEvent struct {
	PlayerId int
}
