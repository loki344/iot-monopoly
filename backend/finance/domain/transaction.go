package financeDomain

import (
	"iot-monopoly/communication"
)

type Transaction struct {
	id          string
	recipientId string
	senderId    string
	amount      int
	accepted    bool
}

func (t *Transaction) SetSenderId(senderId string) {
	t.senderId = senderId
}

func (t *Transaction) Id() string {
	return t.id
}

func (t *Transaction) RecipientId() string {
	return t.recipientId
}

func (t *Transaction) SenderId() string {
	return t.senderId
}

func (t *Transaction) Amount() int {
	return t.amount
}

func (t *Transaction) IsAccepted() bool {
	return t.accepted
}

func (t Transaction) Accept() {

	t.accepted = true
	communication.FireEvent(communication.TRANSACTION_RESOLVED, NewTransactionResolvedEvent(t.id))
}

func NewTransaction(id string, recipientId string, senderId string, amount int) *Transaction {

	transaction := &Transaction{id: id, recipientId: recipientId, senderId: senderId, amount: amount}
	communication.FireEvent(communication.TRANSACTION_CREATED, NewTransactionCreatedEvent(transaction))

	return transaction
}
