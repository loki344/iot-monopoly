package financeDomain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id            string    `json:"id"`
	RecipientId   string    `json:"recipientId"`
	SenderId      string    `json:"senderId"`
	Amount        int       `json:"amount"`
	Accepted      bool      `json:"accepted"`
	ExecutionTime time.Time `json:"executionTime"`
}

func NewTransaction(recipientId string, senderId string, amount int) *Transaction {

	return NewTransactionWithId(uuid.NewString(), recipientId, senderId, amount)
}

func NewTransactionWithId(id string, recipientId string, senderId string, amount int) *Transaction {

	return &Transaction{Id: id, RecipientId: recipientId, SenderId: senderId, Amount: amount, ExecutionTime: time.Time{}}
}

func (transaction *Transaction) IsPending() bool {
	return transaction.ExecutionTime.IsZero()
}
