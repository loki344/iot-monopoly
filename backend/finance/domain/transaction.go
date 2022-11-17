package financeDomain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id            string      `json:"id"`
	RecipientId   string      `json:"recipientId"`
	SenderId      string      `json:"senderId"`
	Amount        int         `json:"amount"`
	Accepted      bool        `json:"accepted"`
	ExecutionTime time.Time   `json:"executionTime"`
	CallbackUrl   string      `json:"callbackUrl"`
	CallbackBody  interface{} `json:"callbackBody"`
}

func NewTransaction(recipientId string, senderId string, amount int) *Transaction {

	return &Transaction{Id: uuid.NewString(), RecipientId: recipientId, SenderId: senderId, Amount: amount, ExecutionTime: time.Time{}}
}

func NewTransactionWithCallback(recipientId string, senderId string, amount int, callbackUrl string, callbackBody interface{}) *Transaction {

	return &Transaction{Id: uuid.NewString(), RecipientId: recipientId, SenderId: senderId, Amount: amount, ExecutionTime: time.Time{}, CallbackUrl: callbackUrl, CallbackBody: callbackBody}
}

func (transaction *Transaction) IsPending() bool {
	return transaction.ExecutionTime.IsZero()
}
