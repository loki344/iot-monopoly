package financeDomain

import (
	"time"
)

type Transaction struct {
	id            string
	recipientId   string
	senderId      string
	amount        uint32
	Accepted      bool
	ExecutionTime time.Time
}

func NewTransaction(id string, recipientId string, senderId string, amount uint32) *Transaction {

	return &Transaction{id: id, recipientId: recipientId, senderId: senderId, amount: amount, ExecutionTime: time.Time{}}
}

func (transaction *Transaction) RecipientId() string {
	return transaction.recipientId
}

func (transaction *Transaction) SenderId() string {
	return transaction.senderId
}

func (transaction *Transaction) Id() string {
	return transaction.id
}

func (transaction *Transaction) IsPending() bool {
	return transaction.ExecutionTime.IsZero()
}

func (transaction *Transaction) Amount() uint32 {
	return transaction.amount
}
