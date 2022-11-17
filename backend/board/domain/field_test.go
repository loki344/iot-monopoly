package boardDomain

import (
	"context"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/config"
	financeDomain "iot-monopoly/finance/domain"
	"testing"
)

func TestPlayerOnOwnerlessFieldFiresBuyQuestionEvent(t *testing.T) {

	config.Init()
	id := uuid.New().String()

	var receivedEvents = 0
	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(*PropertyBuyQuestion)
			assert.Equal(t, id, lapFinishedEvent.PlayerId)
			receivedEvents++
		},
		Matcher: string(eventing.PROPERTY_BUY_QUESTION),
	})

	var tempFinancialDetails = &FinancialDetails{100, 100, 100, Revenue{100, 200, 300, 400, 500, 800}}
	property := NewPropertyField("Property green 2", uuid.NewString(), tempFinancialDetails)
	property.OnPlayerEnter(&Player{id, 0, "Account_Player_1"})

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerOnOwnedFieldFiresTransactionRequestEvent(t *testing.T) {

	config.Init()
	payerId := uuid.New().String()
	ownerId := uuid.New().String()

	var receivedEvents = 0
	const price = 1000
	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transaction := e.Data.(financeDomain.TransactionAddedEvent)
			assert.Equal(t, payerId, transaction.Transaction.SenderId)
			assert.Equal(t, ownerId, transaction.Transaction.RecipientId)
			assert.Equal(t, price, transaction.Transaction.Amount)
			receivedEvents++
		},
		Matcher: string(eventing.TRANSACTION_REQUEST),
	})

	var tempFinancialDetails = &FinancialDetails{100, 100, 100, Revenue{1000, 200, 300, 400, 500, 800}}
	property := NewPropertyField("Property green 2", uuid.NewString(), tempFinancialDetails)
	property.OwnerId = ownerId

	property.OnPlayerEnter(&Player{payerId, 0, "Account_Player_1"})

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerOnEventFieldTriggersEvent(t *testing.T) {

	var eventTriggered = false
	field := EventField{Id: uuid.New().String(), Name: "testField", Event: func(player *Player) {
		eventTriggered = true
	}}
	field.OnPlayerEnter(&Player{Id: uuid.New().String()})
	assert.Equal(t, true, eventTriggered)
}
