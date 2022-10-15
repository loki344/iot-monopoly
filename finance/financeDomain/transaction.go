package financeDomain

import (
	"fmt"
	"iot-monopoly/board"
	"time"
)

type Transaction struct {
	id            string
	recipientId   string
	senderId      string
	amount        int
	Accepted      bool
	executionTime time.Time
}

func NewTransaction(id string, recipientId string, senderId string, amount int) *Transaction {

	return &Transaction{id: id, recipientId: recipientId, senderId: senderId, amount: amount, executionTime: time.Time{}}
}

func (transaction *Transaction) RecipientId() string {
	return transaction.recipientId
}

func (transaction *Transaction) SenderId() string {
	return transaction.senderId
}

func (transaction *Transaction) ExecutionTime() time.Time {
	return transaction.executionTime
}

func (transaction *Transaction) Id() string {
	return transaction.id
}

func (transaction *Transaction) IsPending() bool {
	return transaction.executionTime.IsZero()
}

func (transaction *Transaction) Amount() int {
	return transaction.amount
}

func (transaction *Transaction) Resolve() {
	if !transaction.IsPending() {
		panic(fmt.Sprintf("Transaction %s is already resolved", transaction.id))
	}
	sender := board.GetPlayer(transaction.senderId)
	recipient := board.GetPlayer(transaction.recipientId)

	fmt.Printf("Transferring %d from player %s to player %s\n", transaction.amount, sender.Id, recipient.Id)
	recipient.Balance += transaction.amount
	sender.Balance -= transaction.amount
	transaction.executionTime = time.Now()
}
