package financeAdapter

import (
	"errors"
	"fmt"
	"iot-monopoly/finance/adapter/repository"
	domain "iot-monopoly/finance/domain"
)

func GetAccountById(accountId string) (*domain.Account, error) {

	return repository.FindAccountById(accountId)
}

func AddTransaction(id string, receiverId string, senderId string, amount int) *domain.Transaction {

	transaction := domain.NewTransaction(id, receiverId, senderId, amount)
	repository.CreateTransaction(transaction)

	return transaction
}

func validateTransaction(transaction *domain.Transaction) error {

	if transaction.Accepted {
		return errors.New(fmt.Sprintf("cannot add already accepted transaction, please add only pending pendingTransaction: %s", transaction.Id))
	}
	balance := repository.GetAccountByPlayerId(transaction.SenderId).Balance()
	if balance < transaction.Amount {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func ResolveLatestTransaction(cardId string) error {

	pendingTransaction := repository.GetPendingTransaction()
	if pendingTransaction == nil {
		fmt.Println("No transaction to resolve....")
		return nil
	}

	payerAccount, _ := GetAccountById(cardId)

	if pendingTransaction.SenderId != payerAccount.PlayerId() {
		fmt.Printf("Transaction was meant for cardId %s, but received cardId %s", pendingTransaction.SenderId, cardId)
	}

	err := validateTransaction(pendingTransaction)
	if err != nil {
		return err
	}

	transferFunds(pendingTransaction, payerAccount)
	pendingTransaction.Accept()
	repository.DeleteTransaction(pendingTransaction)

	return nil
}

func transferFunds(pendingTransaction *domain.Transaction, payerAccount *domain.Account) {

	fmt.Printf("Resolving Transaction %s: Transferring %d from senderAccount %s to recipientAccount %s\n", pendingTransaction.Id, pendingTransaction.Amount, pendingTransaction.SenderId, pendingTransaction.RecipientId)
	repository.GetAccountByPlayerId(pendingTransaction.RecipientId).Add(pendingTransaction.Amount)
	payerAccount.Subtract(pendingTransaction.Amount)
}

func GetAccounts() []*AccountDTO {
	accounts := repository.GetAccounts()
	dtos := make([]*AccountDTO, len(accounts))
	for i := range accounts {
		dtos[i] = dtoFromAccount(&accounts[i])
	}
	return dtos
}

type AccountDTO struct {
	Id       string `json:"id"`
	PlayerId string `json:"playerId"`
	Balance  int    `json:"balance"`
}

func dtoFromAccount(account *domain.Account) *AccountDTO {
	return &AccountDTO{account.Id(), account.PlayerId(), account.Balance()}
}
