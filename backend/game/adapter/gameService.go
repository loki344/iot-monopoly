package adapter

import (
	"errors"
	"fmt"
	"iot-monopoly/game/adapter/repository"
	domain "iot-monopoly/game/domain"
)

func NewGame(playerCount int) *domain.Game {
	return repository.SaveGame(domain.NewGame(playerCount))
}

func GetCurrentGame() *domain.Game {
	return repository.GetCurrentGame()
}

func BuyProperty(propertyIndex int, buyerId string) {
	referenceId := repository.GetCurrentGame().BuyProperty(propertyIndex, buyerId)
	transaction := domain.NewTransaction(referenceId, "Bank", GetCurrentGame().GetPlayerById(buyerId).Account().Id(), repository.GetCurrentGame().GetPropertyByIndex(propertyIndex).GetPrice())
	repository.CreateTransaction(transaction)
}

func ResolveCurrentTransaction(cardId string) (*domain.Transaction, error) {
	pendingTransaction := repository.GetPendingTransaction()
	payerAccount := GetCurrentGame().FindAccountById(cardId)

	if pendingTransaction.SenderId() != payerAccount.PlayerId() {
		fmt.Printf("Transaction was meant for cardId %s, but received cardId %s", pendingTransaction.SenderId, cardId)
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
	balance := GetCurrentGame().FindAccountById(transaction.SenderId()).Balance()
	if balance < transaction.Amount() {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func transferFunds(pendingTransaction *domain.Transaction, payerAccount *domain.Account) {

	fmt.Printf("Resolving Transaction %s: Transferring %d from senderAccount %s to recipientAccount %s\n", pendingTransaction.Id, pendingTransaction.Amount, pendingTransaction.SenderId, pendingTransaction.RecipientId)
	GetCurrentGame().FindAccountById(pendingTransaction.RecipientId()).Deposit(pendingTransaction.Amount())
	payerAccount.Pay(pendingTransaction.Amount())
}

func ConfirmCurrentCard() {
	GetCurrentGame().ConfirmCurrentCard()
}
