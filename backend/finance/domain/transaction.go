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

func NewTransactionFromTransactionDTO(transaction *Transaction) *Transaction {
	return &Transaction{Id: uuid.NewString(), RecipientId: transaction.RecipientId, SenderId: transaction.SenderId, Amount: transaction.Amount, ExecutionTime: time.Time{}}
}

func NewTransaction(recipientId string, senderId string, amount int) *Transaction {

	return &Transaction{Id: uuid.NewString(), RecipientId: recipientId, SenderId: senderId, Amount: amount, ExecutionTime: time.Time{}}
}

func (transaction *Transaction) IsPending() bool {
	return transaction.ExecutionTime.IsZero()
}
