package financeDomain

import (
	"iot-monopoly/communication"
)

type Transaction struct {
	Id          string `json:"id"`
	RecipientId string `json:"recipientId"`
	SenderId    string `json:"senderId"`
	Amount      int    `json:"amount"`
	accepted    bool
}

func (t *Transaction) IsAccepted() bool {
	return t.accepted
}

func (t Transaction) Accept() {

	t.accepted = true
	communication.FireEvent(communication.TRANSACTION_RESOLVED, NewTransactionResolvedEvent(t.Id))
}

func NewTransaction(id string, recipientId string, senderId string, amount int) *Transaction {

	transaction := &Transaction{Id: id, RecipientId: recipientId, SenderId: senderId, Amount: amount}
	communication.FireEvent(communication.TRANSACTION_CREATED, NewTransactionCreatedEvent(transaction))

	return transaction
}
