package board

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/eventing"
	"testing"
)

func TestPlayerCanMoveAround(t *testing.T) {

	playerId := StartGame(1)[0].Id

	player := GetPlayer(playerId)
	for i := 0; i < GetFieldsCount(); i++ {
		err := MovePlayer(playerId, i)
		//TODO determine prison fieldindex somehow different
		if i == 12 {
			assert.NoError(t, err)
			//TODO resolve prison index differently
			assert.Equal(t, 4, player.Position)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, i, player.Position)

		}
	}
}

func TestPlayerCannotMoveOutsideBoard(t *testing.T) {

	id := StartGame(1)[0].Id

	errorLowerBound := MovePlayer(id, -1)
	assert.Error(t, errorLowerBound)
	player := GetPlayer(id)
	assert.Equal(t, 0, player.Position)

	errorUpperBound := MovePlayer(id, GetFieldsCount()+1)
	assert.Error(t, errorUpperBound)
	assert.Equal(t, 0, player.Position)
}

func TestLapFiresEvent(t *testing.T) {

	id := StartGame(1)[0].Id

	lapFinishedEventHandler := eventing.ListenRequestStream(eventing.LAP_FINISHED)

	var receivedEvents = 0
	lapFinishedEventHandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(domain.LapFinishedEvent)
			assert.Equal(t, id, lapFinishedEvent.PlayerId)
			receivedEvents++
		},
		func(err error) {
			fmt.Println(err)
		})

	MovePlayer(id, 15)
	MovePlayer(id, 1)

	assert.Equal(t, 1, receivedEvents)
}
