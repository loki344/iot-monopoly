package finance

import (
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/config"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	config.Init()
	StartEventListeners()

	players, _ := board.StartGame(1)
	id := players[0].Id
	eventing.FireEvent(eventing.LAP_FINISHED, &boardDomain.LapFinishedEvent{PlayerId: id})
	assert.Equal(t, 1100, getAccountByPlayerId(id).Balance)
}
