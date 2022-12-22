package financeDomain

import eventingDomain "iot-monopoly/communication/domain"

type TransactionCreatedEvent struct {
	eventingDomain.BaseEvent
	Transaction *Transaction
}

type TransactionResolvedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
}

func NewTransactionCreatedEvent(transaction *Transaction) TransactionCreatedEvent {

	return TransactionCreatedEvent{BaseEvent: eventingDomain.EventType(&TransactionCreatedEvent{}), Transaction: transaction}
}

func NewTransactionResolvedEvent(transactionId string) TransactionResolvedEvent {

	return TransactionResolvedEvent{BaseEvent: eventingDomain.EventType(&TransactionResolvedEvent{}), TransactionId: transactionId}
}
