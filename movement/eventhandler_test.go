package movement

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
	"testing"
)

func TestSensorEventsMovePlayer(t *testing.T) {

	StartEventHandler()

	id := uuid.New().String()
	board.StartGame([]boardDomain.Player{{id, 0, 1000}})
	eventing.FireEvent(eventing.EXTERNAL, eventing.SensorEvent{id, 3})
	assert.Equal(t, 3, board.GetPlayer(id).Position)
}
