package propertyDomain

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
			lapFinishedEvent := e.Data.(*PlayerOnFieldEvent)
			assert.Equal(t, id, lapFinishedEvent.PlayerId)
			receivedEvents++
		},
		Matcher: string(communication.PLAYER_ON_UNOWNED_FIELD),
	})

	var tempFinancialDetails = &FinancialDetails{100, 100, 100, Revenue{100, 200, 300, 400, 500, 800}}
	property := NewPropertyField(BaseFieldInformation{"Property green 2", "2"}, tempFinancialDetails)
	property.OnPlayerEnter(id)

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerOnEventFieldTriggersEvent(t *testing.T) {

	var eventTriggered = false
	field := EventField{BaseFieldInformation: BaseFieldInformation{Id: uuid.New().String(), Name: "testField"}, Event: func(playerId string) {
		eventTriggered = true
	}}
	field.OnPlayerEnter(uuid.NewString())
	assert.Equal(t, true, eventTriggered)
}
