package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/communication"
	"iot-monopoly/communication/config"
	"testing"
)

func TestPlayerOnOwnerlessFieldFiresBuyQuestionEvent(t *testing.T) {

	config.Init()
	id := uuid.New().String()

	var receivedEvents = 0
	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(*PlayerOnUnownedFieldEvent)
			assert.Equal(t, id, lapFinishedEvent.PlayerId)
			receivedEvents++
		},
		Matcher: string(communication.PLAYER_ON_UNOWNED_FIELD),
	})

	var tempFinancialDetails = &FinancialDetails{100, 100, 100, Revenue{100, 200, 300, 400, 500, 800}}
	property := NewPropertyField("Property green 2", "2", tempFinancialDetails)
	property.OnPlayerEnter(id)

	assert.Equal(t, 1, receivedEvents)
}
