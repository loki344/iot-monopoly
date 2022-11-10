package finance

import (
	"errors"
	"fmt"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
	"iot-monopoly/finance/domain"
	"time"
)

var transactions []financeDomain.Transaction

func AddTransaction(transaction financeDomain.Transaction) (*financeDomain.Transaction, error) {

	err := validateTransaction(transaction)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Adding transaction %s to pending transactions\n", transaction.Id)
	transactions = append(transactions, transaction)

	eventing.FireEvent(eventing.TRANSACTION_ADDED, financeDomain.NewTransactionRequest(&transaction))

	return &transaction, nil
}

func validateTransaction(transaction financeDomain.Transaction) error {

	if !transaction.IsPending() {
		return errors.New(fmt.Sprintf("cannot add non-pending transaction, please add only pending transactions: %s", transaction.Id))
	}
	sender := board.GetPlayer(transaction.SenderId)
	if sender.Balance < transaction.Amount {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, sender.Balance, transaction.Amount))
	}
	return nil
}

func ResolveTransaction(id string) {
	transaction := GetTransaction(id)
	if !transaction.IsPending() {
		panic(fmt.Sprintf("Transaction %s is already resolved", transaction.Id))
	}
	sender := board.GetPlayer(transaction.SenderId)
	recipient := board.GetPlayer(transaction.RecipientId)

	fmt.Printf("Transferring %d from player %s to player %s\n", transaction.Amount, sender.Id, recipient.Id)
	recipient.Balance += transaction.Amount
	sender.Balance -= transaction.Amount

	transactions = remove(transactions, transaction)
	transaction.ExecutionTime = time.Now()
	transactions = append(transactions, transaction)
}

func remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func GetTransaction(id string) financeDomain.Transaction {

	for _, transaction := range transactions {
		if transaction.Id == id {
			return transaction
		}
	}

	panic(fmt.Sprintf("no transaction found with id %s", id))
}

func getPendingTransactions() []financeDomain.Transaction {
	return transactions
}
