package propertyAdapter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/communication/config"
	"testing"
)

func TestBuyProperty(t *testing.T) {

	config.Init()
	StartEventListeners()
	initFields()
	playerId := uuid.NewString()

	propertyId := "2"
	transactionId := BuyProperty(propertyId, playerId)

	transferOwnerShip(transactionId)

	property := *getPropertyById(propertyId)
	assert.Equal(t, playerId, property.OwnerId)
}
