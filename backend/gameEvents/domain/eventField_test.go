package gameEventsDomain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerOnEventFieldTriggersEvent(t *testing.T) {

	var eventTriggered = false
	field := EventField{id: uuid.New().String(), name: "testField", event: func(playerId string) {
		eventTriggered = true
	}}
	field.OnPlayerEnter(uuid.NewString())
	assert.Equal(t, true, eventTriggered)
}
