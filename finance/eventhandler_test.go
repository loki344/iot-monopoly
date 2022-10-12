package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	StartEventHandler()

	id := uuid.New().String()
	board.StartGame([]boardDomain.Player{{id, 0, 1000}})
	eventing.FireEvent(eventing.LAP_FINISHED, boardDomain.LapFinishedEvent{id})
	assert.Equal(t, 1100, board.GetPlayer(id).Balance)
}
