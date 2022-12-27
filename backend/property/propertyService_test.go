package property

import (
	"github.com/stretchr/testify/assert"
	"iot-monopoly/communication/config"
	"iot-monopoly/player"
	"testing"
)

func TestBuyProperty(t *testing.T) {

	config.Init()
	StartEventListeners()
	initFields()
	players, _ := player.Init(1)
	playerId := players[0].Id

	propertyId := "2"
	transactionId := BuyProperty(propertyId, playerId)

	transferOwnerShip(transactionId)

	property := *getPropertyById(propertyId)
	assert.Equal(t, playerId, property.OwnerId)
}
