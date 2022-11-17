package finance

import (
	"errors"
	"fmt"
	"iot-monopoly/eventing"
	"iot-monopoly/finance/domain"
	"time"
)

var transactions []financeDomain.Transaction
var defaultAccounts = []financeDomain.Account{
	{"Account_Player_1", "Player_1", 1_000},
	{"Account_Player_2", "Player_2", 1_000},
	{"Account_Player_3", "Player_3", 1_000},
	{"Account_Player_4", "Player_4", 1_000},
	{"Bank", "Bank", 100_000_000},
}
var accounts = defaultAccounts

func InitAccounts() {
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

func addToAccount(accountId string, amount int) {
	fmt.Printf("Adding %d to account %s\n", amount, accountId)
	account, _ := GetAccountById(accountId)
	account.Balance += amount
}

func removeFromAccount(accountId string, amount int) {
	fmt.Printf("Removing %d from account %s\n", amount, accountId)
	account, _ := GetAccountById(accountId)
	account.Balance -= amount

}

func AddTransaction(transaction *financeDomain.Transaction) (*financeDomain.Transaction, error) {

	err := validateTransaction(transaction)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Adding transaction %s to pending transactions\n", transaction.Id)
	transactions = append(transactions, *transaction)

	eventing.FireEvent(eventing.TRANSACTION_REQUEST, financeDomain.NewTransactionRequest(transaction))

	return transaction, nil
}

func validateTransaction(transaction *financeDomain.Transaction) error {

	if !transaction.IsPending() {
		return errors.New(fmt.Sprintf("cannot add non-pending transaction, please add only pending transactions: %s", transaction.Id))
	}
	balance := getAccountByPlayerId(transaction.SenderId).Balance
	if balance < transaction.Amount {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func ResolveTransaction(id string) {
	transaction := GetTransaction(id)
	if !transaction.IsPending() {
		panic(fmt.Sprintf("Transaction %s is already resolved", transaction.Id))
	}

	fmt.Printf("Resolving Transaction %s: Transferring %d from account %s to account %s\n", transaction.Id, transaction.Amount, transaction.SenderId, transaction.RecipientId)
	addToAccount(getAccountByPlayerId(transaction.RecipientId).Id, transaction.Amount)
	removeFromAccount(getAccountByPlayerId(transaction.SenderId).Id, transaction.Amount)

	transaction.ExecutionTime = time.Now()
	eventing.FireEvent(eventing.TRANSACTION_RESOLVED, financeDomain.NewTransactionResolvedEvent(transaction.Id))
}

func GetAccounts() []financeDomain.Account {
	return accounts
}

func GetTransaction(id string) *financeDomain.Transaction {

	for i := range transactions {
		if transactions[i].Id == id {
			return &transactions[i]
		}
	}

	panic(fmt.Sprintf("no transaction found with id %s", id))
}

func getPendingTransactions() []financeDomain.Transaction {
	return transactions
}
