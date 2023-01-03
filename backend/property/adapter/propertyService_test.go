package adapter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/eventing/config"
	"iot-monopoly/property/adapter/repository"
	"testing"
)

func TestBuyProperty(t *testing.T) {

	config.Init()
	StartEventListeners()
	repository.InitFields()
	playerId := uuid.NewString()

	propertyId := "2"
	transactionId := BuyProperty(propertyId, playerId)

	transferOwnerShip(transactionId)

	property := *getPropertyById(propertyId)
	assert.Equal(t, playerId, property.OwnerId)
}
