package financeDomain

import eventingDomain "iot-monopoly/eventing/domain"

type TransactionRequested struct {
	eventingDomain.BaseEvent
	Transaction *Transaction
}

func NewTransactionRequest(transaction *Transaction) TransactionRequested {

	return TransactionRequested{BaseEvent: eventingDomain.EventType(&TransactionRequested{}), Transaction: transaction}
}
