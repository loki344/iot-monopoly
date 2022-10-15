package board

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/domain"
	"testing"
)

func TestPlayerCanMoveAround(t *testing.T) {

	id := uuid.New().String()
	StartGame([]boardDomain.Player{{id, 0, 1000}})

	player := GetPlayer(id)
	for i := 0; i < GetFieldsCount(); i++ {
		err := MovePlayer(id, i)
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

	id := uuid.New().String()
	StartGame([]boardDomain.Player{{id, 0, 1000}})

	errorLowerBound := MovePlayer(id, -1)
	assert.Error(t, errorLowerBound)
	player := GetPlayer(id)
	assert.Equal(t, 0, player.Position)

	errorUpperBound := MovePlayer(id, GetFieldsCount()+1)
	assert.Error(t, errorUpperBound)
	assert.Equal(t, 0, player.Position)
}

func TestLapFiresEvent(t *testing.T) {

	id := uuid.New().String()
	StartGame([]boardDomain.Player{{id, 0, 1000}})

	lapFinishedEventHandler := eventing.ListenRequestStream(eventing.LAP_FINISHED)

	var receivedEvents = 0
	lapFinishedEventHandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(eventingDomain.LapFinishedEvent)
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
