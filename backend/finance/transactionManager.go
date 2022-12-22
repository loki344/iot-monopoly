package finance

import (
	"errors"
	"fmt"
	"iot-monopoly/communication"
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

	fmt.Printf("Adding transaction %s to pending pendingTransaction\n", transaction.Id)
	pendingTransaction = transaction

	communication.FireEvent(communication.TRANSACTION_CREATED, financeDomain.NewTransactionCreatedEvent(transaction))

	return transaction, nil
}

func validateTransaction(transaction *financeDomain.Transaction) error {

	if transaction.Accepted {
		return errors.New(fmt.Sprintf("cannot add already accepted transaction, please add only pending pendingTransaction: %s", transaction.Id))
	}
	balance := getAccountByPlayerId(transaction.SenderId).Balance
	if balance < transaction.Amount {
		return errors.New(fmt.Sprintf("Player %s has insufficient balance for transaction %s. Balance: %d, amount: %d", transaction.SenderId, transaction.Id, balance, transaction.Amount))
	}
	return nil
}

func ResolveLatestTransaction(senderId string) {

	if pendingTransaction == nil {
		fmt.Println("No transaction to resolve....")
		return
	}

	if pendingTransaction.SenderId != senderId {
		fmt.Printf("Transaction was meant for senderId %s, but received senderId %s", pendingTransaction.SenderId, senderId)
		pendingTransaction.SenderId = senderId
		err := validateTransaction(pendingTransaction)
		if err != nil {
			fmt.Println(err)
			//TODO improve
			panic("pendingTransaction invalid")
		}
	}

	fmt.Printf("Resolving Transaction %s: Transferring %d from account %s to account %s\n", pendingTransaction.Id, pendingTransaction.Amount, pendingTransaction.SenderId, pendingTransaction.RecipientId)
	addToAccount(getAccountByPlayerId(pendingTransaction.RecipientId).Id, pendingTransaction.Amount)
	removeFromAccount(getAccountByPlayerId(pendingTransaction.SenderId).Id, pendingTransaction.Amount)

	pendingTransaction.Accepted = true
	communication.FireEvent(communication.TRANSACTION_RESOLVED, financeDomain.NewTransactionResolvedEvent(pendingTransaction.Id))
	pendingTransaction = nil
}

func GetAccounts() []financeDomain.Account {
	return accounts
}
