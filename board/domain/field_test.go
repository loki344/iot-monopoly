package boardDomain

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/domain"
	"testing"
)

func TestPlayerOnOwnerlessFieldFiresBuyQuestionEvent(t *testing.T) {

	propertyBuyQuestionEventHandler := eventing.ListenRequestStream(eventing.PROPERTY_BUY_QUESTION)
	id := uuid.New().String()

	var receivedEvents = 0
	propertyBuyQuestionEventHandler.Handle(
		func(msg *model.Message) {
			buyQuestionEvent := msg.Payload.(eventingDomain.PropertyBuyQuestion)
			assert.Equal(t, id, buyQuestionEvent.PlayerId)
			receivedEvents++
		},
		func(err error) {
			fmt.Println(err)
			t.Fail()
		})

	property := PropertyField{"TestProperty", uuid.New().String(), 1000, ""}
	property.OnPlayerEnter(&Player{id, 0, 1000})

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerOnOwnedFieldFiresTransactionRequestEvent(t *testing.T) {

	propertyBuyQuestionEventHandler := eventing.ListenRequestStream(eventing.TRANSACTION_REQUESTED)
	payerId := uuid.New().String()
	ownerId := uuid.New().String()

	var receivedEvents = 0
	const price = 1000
	propertyBuyQuestionEventHandler.Handle(
		func(msg *model.Message) {
			transactionRequest := msg.Payload.(eventingDomain.TransactionRequested)
			assert.Equal(t, payerId, transactionRequest.SenderId())
			assert.Equal(t, ownerId, transactionRequest.RecipientId())
			assert.Equal(t, price, transactionRequest.Amount())
			receivedEvents++
		},
		func(err error) {
			fmt.Println(err)
			t.Fail()
		})

	property := PropertyField{"TestProperty", uuid.New().String(), price, ownerId}
	property.OnPlayerEnter(&Player{payerId, 0, 0})

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
