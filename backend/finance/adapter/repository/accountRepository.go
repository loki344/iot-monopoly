package repository

import (
	"errors"
	"fmt"
	domain "iot-monopoly/finance/domain"
)

var defaultAccounts = []domain.Account{
	*domain.NewAccount("Account_Player_1", "Player_1", 1_000),
	*domain.NewAccount("Account_Player_2", "Player_2", 1_000),
	*domain.NewAccount("Account_Player_3", "Player_3", 1_000),
	*domain.NewAccount("Account_Player_4", "Player_4", 1_000),
	*domain.NewAccount("Bank", "Bank", 100_000_000),
}
var accounts = defaultAccounts

func InitAccounts() {
	accounts = []domain.Account{
		*domain.NewAccount("Account_Player_1", "Player_1", 1_000),
		*domain.NewAccount("Account_Player_2", "Player_2", 1_000),
		*domain.NewAccount("Account_Player_3", "Player_3", 1_000),
		*domain.NewAccount("Account_Player_4", "Player_4", 1_000),
		*domain.NewAccount("Bank", "Bank", 100_000_000),
	}
}

func FindAccountById(accountId string) (*domain.Account, error) {

	for i := range accounts {
		if accounts[i].Id() == accountId {
			return &accounts[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No account with id %s found", accountId))
}

func GetAccountByPlayerId(playerId string) *domain.Account {
	for i := range accounts {
		if accounts[i].PlayerId() == playerId {
			return &accounts[i]
		}
	}
	panic(fmt.Sprintf("No account for playerId %s found", playerId))
}

func GetAccounts() []domain.Account {
	return accounts
}
