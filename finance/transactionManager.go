package finance

import (
	"errors"
	"fmt"
	"iot-monopoly/board"
	"iot-monopoly/finance/financeDomain"
)

var transactions []financeDomain.Transaction

func AddTransaction(transaction financeDomain.Transaction) error {

	err := validateTransaction(transaction)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Adding transaction %s to pending transactions\n", transaction.Id())
	transactions = append(transactions, transaction)
	return nil
}

func validateTransaction(transaction financeDomain.Transaction) error {

	if !transaction.IsPending() {
		return errors.New(fmt.Sprintf("cannot add non-pending transaction, please add only pending transactions: %s", transaction.Id()))
	}
	sender := board.GetPlayer(transaction.SenderId())
	if sender.Balance < transaction.Amount() {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId(), transaction.Id(), sender.Balance, transaction.Amount()))
	}
	return nil
}

func GetTransaction(id string) *financeDomain.Transaction {

	for _, transaction := range transactions {
		if transaction.Id() == id {
			return &transaction
		}
	}

	panic(fmt.Sprintf("no transaction found with id %s", id))
}

func getPendingTransactions() []financeDomain.Transaction {
	return transactions
}
