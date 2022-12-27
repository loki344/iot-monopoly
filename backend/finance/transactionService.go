package finance

import (
	"errors"
	"fmt"
	"iot-monopoly/finance/domain"
)

var pendingTransaction *financeDomain.Transaction
var defaultAccounts = []financeDomain.Account{
	{"Account_Player_1", "Player_1", 1_000},
	{"Account_Player_2", "Player_2", 1_000},
	{"Account_Player_3", "Player_3", 1_000},
	{"Account_Player_4", "Player_4", 1_000},
	{"Bank", "Bank", 100_000_000},
}
var accounts = defaultAccounts

func initAccounts() {
	accounts = []financeDomain.Account{
		{"Account_Player_1", "Player_1", 1_000},
		{"Account_Player_2", "Player_2", 1_000},
		{"Account_Player_3", "Player_3", 1_000},
		{"Account_Player_4", "Player_4", 1_000},
		{"Bank", "Bank", 100_000_000},
	}
}

func GetAccountById(accountId string) (*financeDomain.Account, error) {

	for i := range accounts {
		if accounts[i].Id == accountId {
			return &accounts[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No account with id %s found", accountId))
}

func getAccountByPlayerId(playerId string) *financeDomain.Account {
	for i := range accounts {
		if accounts[i].PlayerId == playerId {
			return &accounts[i]
		}
	}
	panic(fmt.Sprintf("No account for playerId %s found", playerId))
}

func AddTransaction(id string, receiverId string, senderId string, amount int) *financeDomain.Transaction {

	transaction := financeDomain.NewTransaction(id, receiverId, senderId, amount)

	fmt.Printf("Adding transaction %s to pending pendingTransaction\n", transaction.Id)
	pendingTransaction = transaction

	return transaction
}

func validateTransaction(transaction *financeDomain.Transaction) error {

	if transaction.IsAccepted() {
		return errors.New(fmt.Sprintf("cannot add already accepted transaction, please add only pending pendingTransaction: %s", transaction.Id))
	}
	balance := getAccountByPlayerId(transaction.SenderId).Balance
	if balance < transaction.Amount {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func ResolveLatestTransaction(senderId string) error {

	if pendingTransaction == nil {
		fmt.Println("No transaction to resolve....")
		return nil
	}

	if pendingTransaction.SenderId != senderId {
		fmt.Printf("Transaction was meant for senderId %s, but received senderId %s", pendingTransaction.SenderId, senderId)
		pendingTransaction.SenderId = senderId
	}

	err := validateTransaction(pendingTransaction)
	if err != nil {
		return err
	}

	fmt.Printf("Resolving Transaction %s: Transferring %d from account %s to account %s\n", pendingTransaction.Id, pendingTransaction.Amount, pendingTransaction.SenderId, pendingTransaction.RecipientId)
	getAccountByPlayerId(pendingTransaction.RecipientId).Add(pendingTransaction.Amount)
	getAccountByPlayerId(pendingTransaction.SenderId).Subtract(pendingTransaction.Amount)

	pendingTransaction.Accept()
	pendingTransaction = nil
	return nil
}

func GetAccounts() []financeDomain.Account {
	return accounts
}
