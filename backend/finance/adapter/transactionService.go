package financeAdapter

import (
	"errors"
	"fmt"
	domain "iot-monopoly/finance/domain"
)

var pendingTransaction *domain.Transaction
var defaultAccounts = []domain.Account{
	*domain.NewAccount("Account_Player_1", "Player_1", 1_000),
	*domain.NewAccount("Account_Player_2", "Player_2", 1_000),
	*domain.NewAccount("Account_Player_3", "Player_3", 1_000),
	*domain.NewAccount("Account_Player_4", "Player_4", 1_000),
	*domain.NewAccount("Bank", "Bank", 100_000_000),
}
var accounts = defaultAccounts

func initAccounts() {
	accounts = []domain.Account{
		*domain.NewAccount("Account_Player_1", "Player_1", 1_000),
		*domain.NewAccount("Account_Player_2", "Player_2", 1_000),
		*domain.NewAccount("Account_Player_3", "Player_3", 1_000),
		*domain.NewAccount("Account_Player_4", "Player_4", 1_000),
		*domain.NewAccount("Bank", "Bank", 100_000_000),
	}
}

func GetAccountById(accountId string) (*domain.Account, error) {

	for i := range accounts {
		if accounts[i].Id() == accountId {
			return &accounts[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No account with id %s found", accountId))
}

func getAccountByPlayerId(playerId string) *domain.Account {
	for i := range accounts {
		if accounts[i].PlayerId() == playerId {
			return &accounts[i]
		}
	}
	panic(fmt.Sprintf("No account for playerId %s found", playerId))
}

func AddTransaction(id string, receiverId string, senderId string, amount int) *domain.Transaction {

	transaction := domain.NewTransaction(id, receiverId, senderId, amount)

	fmt.Printf("Adding transaction %s to pending pendingTransaction\n", transaction.Id)
	pendingTransaction = transaction

	return transaction
}

func validateTransaction(transaction *domain.Transaction) error {

	if transaction.Accepted {
		return errors.New(fmt.Sprintf("cannot add already accepted transaction, please add only pending pendingTransaction: %s", transaction.Id))
	}
	balance := getAccountByPlayerId(transaction.SenderId).Balance()
	if balance < transaction.Amount {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func ResolveLatestTransaction(cardId string) error {

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

	fmt.Printf("Resolving Transaction %s: Transferring %d from senderAccount %s to recipientAccount %s\n", pendingTransaction.Id, pendingTransaction.Amount, pendingTransaction.SenderId, pendingTransaction.RecipientId)
	getAccountByPlayerId(pendingTransaction.RecipientId).Add(pendingTransaction.Amount)
	payerAccount.Subtract(pendingTransaction.Amount)

	pendingTransaction.Accept()
	pendingTransaction = nil
	return nil
}

func GetAccounts() []*AccountDTO {
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
