package adapter

import (
	"errors"
	"fmt"
	"iot-monopoly/game/adapter/repository"
	"iot-monopoly/game/domain"
)

func ResolveCurrentTransaction(accountId string) (*domain.Transaction, error) {
	pendingTransaction := repository.GetPendingTransaction()
	payerAccount := repository.FindAccountById(accountId)

	if pendingTransaction.SenderId() != payerAccount.Id() {
		fmt.Printf("Transaction was meant for accountId %s, but received accountId %s", pendingTransaction.SenderId(), accountId)
	}

	err := validateTransaction(pendingTransaction)
	if err != nil {
		return nil, err
	}

	transferFunds(pendingTransaction, payerAccount)
	pendingTransaction.Accept()
	repository.DeleteTransaction(pendingTransaction)
	return pendingTransaction, nil
}

func validateTransaction(transaction *domain.Transaction) error {

	if transaction.Accepted() {
		return errors.New(fmt.Sprintf("cannot add already accepted transaction, please add only pending pendingTransaction: %s", transaction.Id))
	}
	balance := repository.FindAccountById(transaction.SenderId()).Balance()
	if balance < transaction.Amount() {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func transferFunds(pendingTransaction *domain.Transaction, payerAccount *domain.Account) {

	fmt.Printf("Resolving Transaction %s: Transferring %d from senderAccount %s to recipientAccount %s\n", pendingTransaction.Id, pendingTransaction.Amount, pendingTransaction.SenderId, pendingTransaction.RecipientId)
	repository.FindAccountById(pendingTransaction.RecipientId()).Deposit(pendingTransaction.Amount())
	payerAccount.Pay(pendingTransaction.Amount())
}
