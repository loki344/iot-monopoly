package financeDomain

import eventingDomain "iot-monopoly/eventing/domain"

type TransactionAddedEvent struct {
	eventingDomain.BaseEvent
	Transaction *Transaction
}

func NewTransactionRequest(transaction *Transaction) TransactionAddedEvent {

	return TransactionAddedEvent{BaseEvent: eventingDomain.EventType(&TransactionAddedEvent{}), Transaction: transaction}
}
