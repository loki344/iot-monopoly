package domain

import (
	eventingDomain "iot-monopoly/eventing/domain"
)

type TransactionCreatedEvent struct {
	eventingDomain.BaseEvent
	RecipientId string
	SenderId    string
	Amount      int
}

func NewTransactionCreatedEvent(recipientId string, senderId string, amount int) TransactionCreatedEvent {

	return TransactionCreatedEvent{BaseEvent: eventingDomain.EventType(&TransactionCreatedEvent{}), RecipientId: recipientId, SenderId: senderId, Amount: amount}
}
