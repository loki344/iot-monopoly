package domain

import (
	"iot-monopoly/eventing"
)

type Transaction struct {
	id          string
	recipientId string
	senderId    string
	amount      int
	accepted    bool
}

func (t Transaction) RecipientId() string {
	return t.recipientId
}

func (t Transaction) SenderId() string {
	return t.senderId
}

func (t Transaction) Amount() int {
	return t.amount
}

func (t Transaction) Accepted() bool {
	return t.accepted
}

func (t Transaction) Id() string {
	return t.id
}

func (t Transaction) Accept() {

	t.accepted = true
	eventing.FireEvent(eventing.TRANSACTION_RESOLVED, newTransactionResolvedEvent(t.id))
}

func NewTransaction(id string, recipientId string, senderId string, amount int) *Transaction {

	transaction := &Transaction{id: id, recipientId: recipientId, senderId: senderId, amount: amount}
	eventing.FireEvent(eventing.TRANSACTION_CREATED, newTransactionCreatedEvent(transaction.recipientId, transaction.senderId, transaction.amount))

	return transaction
}
