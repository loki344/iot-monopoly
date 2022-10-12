package board

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board/boardDomain"
	"testing"
)

func TestPlayerCanMoveAround(t *testing.T) {

	id := uuid.New().String()
	StartGame([]boardDomain.Player{{id, 0, 1000}})

	player := GetPlayer(id)
	for i := 0; i < len(Fields); i++ {
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

	errorUpperBound := MovePlayer(id, len(Fields)+1)
	assert.Error(t, errorUpperBound)
	assert.Equal(t, 0, player.Position)
}
