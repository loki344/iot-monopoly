package domain

import (
	"github.com/stretchr/testify/assert"
	"iot-monopoly/eventing/config"
	"testing"
)

func TestPlayerMoveAround(t *testing.T) {

	config.Init()

	game := NewGame(2)
	players := game.Players()
	assert.Equal(t, 1, players[0].Position())
	game.MovePlayer(game.players[0].Id(), 2)
	assert.Equal(t, 2, players[0].Position())
}

func TestLapFinishAddsMoney(t *testing.T) {

	config.Init()

	game := NewGame(2)
	players := game.Players()
	game.MovePlayer(game.players[0].Id(), 6)
	game.MovePlayer(game.players[0].Id(), 12)
	game.MovePlayer(game.players[0].Id(), 1)
	assert.Equal(t, 1_100, players[0].Account().Balance())
}

func TestCurrentPlayerIndex(t *testing.T) {
	config.Init()

	game := NewGame(2)
	game.MovePlayer(game.players[0].Id(), 6)
	assert.Equal(t, 1, game.currentPlayerIndex)

}
