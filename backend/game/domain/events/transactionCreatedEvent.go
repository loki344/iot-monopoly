package events

import (
	eventingDomain "iot-monopoly/eventing/domain"
	"iot-monopoly/game/domain"
)

type TransactionCreatedEvent struct {
	eventingDomain.BaseEvent
	Transaction *domain.Transaction
}

func NewTransactionCreatedEvent(transaction *domain.Transaction) TransactionCreatedEvent {

	return TransactionCreatedEvent{BaseEvent: eventingDomain.EventType(&TransactionCreatedEvent{}), Transaction: transaction}
}
