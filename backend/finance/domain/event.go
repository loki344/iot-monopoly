package financeDomain

import eventingDomain "iot-monopoly/eventing/domain"

type TransactionRequested struct {
	eventingDomain.BaseEvent
	id          string
	recipientId string
	senderId    string
	amount      uint32
}

func (t TransactionRequested) Id() string {
	return t.id
}

func (t TransactionRequested) RecipientId() string {
	return t.recipientId
}

func (t TransactionRequested) SenderId() string {
	return t.senderId
}

func (t TransactionRequested) Amount() uint32 {
	return t.amount
}

func NewTransactionRequest(id string, recipientId string, senderId string, amount uint32) TransactionRequested {

	if amount <= 0 {
		panic("amount has to be greater than 0")
	}
	return TransactionRequested{BaseEvent: eventingDomain.EventType(&TransactionRequested{}), id: id, recipientId: recipientId, senderId: senderId, amount: amount}
}
