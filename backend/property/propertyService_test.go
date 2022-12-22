package property

import (
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/communication/config"
	"testing"
)

func TestBuyProperty(t *testing.T) {

	config.Init()
	StartEventListeners()
	players, _ := board.StartGame(1)
	playerId := players[0].Id

	propertyId := "2"
	transactionId := BuyProperty(propertyId, playerId)

	transferOwnerShip(transactionId)

	property := *GetPropertyById(propertyId)
	assert.Equal(t, playerId, property.OwnerId)
}
