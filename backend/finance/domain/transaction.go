package financeDomain

import (
	"iot-monopoly/eventing"
)

type Transaction struct {
	Id          string `json:"id"`
	RecipientId string `json:"recipientId"`
	SenderId    string `json:"senderId"`
	Amount      int    `json:"amount"`
	Accepted    bool   `json:"accepted"`
}

func (t Transaction) Accept() {

	t.Accepted = true
	eventing.FireEvent(eventing.TRANSACTION_RESOLVED, NewTransactionResolvedEvent(t.Id))
}

func NewTransaction(id string, recipientId string, senderId string, amount int) *Transaction {

	transaction := &Transaction{Id: id, RecipientId: recipientId, SenderId: senderId, Amount: amount}
	eventing.FireEvent(eventing.TRANSACTION_CREATED, NewTransactionCreatedEvent(transaction))

	return transaction
}
