package board

import (
	"context"
	"github.com/mustafaturan/bus/v3"
	"github.com/stretchr/testify/assert"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/communication"
	"iot-monopoly/communication/config"
	"testing"
)

func TestPlayerCanMoveAround(t *testing.T) {

	config.Init()
	players, _ := StartGame(1)
	playerId := players[0].Id
	player := GetPlayer(playerId)
	// TODO get total number count 16
	for i := 1; i < 16; i++ {
		err := MovePlayer(playerId, i)
		//TODO determine prison fieldindex somehow different
		if i == 13 {
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

	players, _ := StartGame(1)
	id := players[0].Id

	player := GetPlayer(id)

	errorUpperBound := MovePlayer(id, 17)
	assert.Error(t, errorUpperBound)
	assert.Equal(t, 0, player.Position)
}

func TestLapFiresEvent(t *testing.T) {

	config.Init()
	players, _ := StartGame(1)
	id := players[0].Id

	var receivedEvents = 0
	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(*boardDomain.LapFinishedEvent)
			assert.Equal(t, id, lapFinishedEvent.PlayerId)
			receivedEvents++
		},
		Matcher: string(communication.LAP_FINISHED),
	})

	MovePlayer(id, 15)
	MovePlayer(id, 1)

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerTurn(t *testing.T) {

	config.Init()
	players, _ := StartGame(4)

	assert.Equal(t, players[0].Id, currentPlayer.Id)
	MovePlayer(players[0].Id, 2)

	assert.Equal(t, players[1].Id, currentPlayer.Id)
	MovePlayer(players[1].Id, 2)

	assert.Equal(t, players[2].Id, currentPlayer.Id)
	MovePlayer(players[2].Id, 2)

	assert.Equal(t, players[3].Id, currentPlayer.Id)
	MovePlayer(players[3].Id, 2)

	assert.Equal(t, players[0].Id, currentPlayer.Id)
	MovePlayer(players[0].Id, 2)

}
