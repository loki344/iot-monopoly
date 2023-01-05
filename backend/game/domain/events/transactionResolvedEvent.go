package events

import eventingDomain "iot-monopoly/eventing/domain"

type TransactionResolvedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
}

func NewTransactionResolvedEvent(transactionId string) TransactionResolvedEvent {

	return TransactionResolvedEvent{BaseEvent: eventingDomain.EventType(&TransactionResolvedEvent{}), TransactionId: transactionId}
}
