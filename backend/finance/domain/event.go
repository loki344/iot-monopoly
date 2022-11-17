package financeDomain

import eventingDomain "iot-monopoly/eventing/domain"

type TransactionAddedEvent struct {
	eventingDomain.BaseEvent
	Transaction *Transaction
}

type TransactionResolvedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
}

func NewTransactionRequest(transaction *Transaction) TransactionAddedEvent {

	return TransactionAddedEvent{BaseEvent: eventingDomain.EventType(&TransactionAddedEvent{}), Transaction: transaction}
}

func NewTransactionResolvedEvent(transactionId string) TransactionResolvedEvent {

	return TransactionResolvedEvent{BaseEvent: eventingDomain.EventType(&TransactionAddedEvent{}), TransactionId: transactionId}
}
