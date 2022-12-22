package finance

import (
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/communication/config"
	"iot-monopoly/finance/domain"
	"testing"
)

func TestTransactionWithInsufficientBalance(t *testing.T) {
	config.Init()
	InitAccounts()
	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	amount := 1_000

	_, err := AddTransaction(financeDomain.NewTransaction(recipientId, senderId, amount))
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidTransaction(t *testing.T) {
	config.Init()
	InitAccounts()

	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	amount := 1_000

	transaction, err := AddTransaction(financeDomain.NewTransaction(recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, recipientId, transaction.RecipientId)
	assert.Equal(t, senderId, transaction.SenderId)
	assert.Equal(t, amount, transaction.Amount)
}

func TestResolveTransactionChangesBalance(t *testing.T) {
	config.Init()
	InitAccounts()

	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	amount := 1_000

	_, err := AddTransaction(financeDomain.NewTransaction(recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}

	ResolveLatestTransaction(senderId)

	assert.Equal(t, 0, getAccountByPlayerId(senderId).Balance)
	assert.Equal(t, 2_000, getAccountByPlayerId(recipientId).Balance)
}

func TestTransactionWithChangedSenderId(t *testing.T) {

	InitAccounts()
	config.Init()

	players, _ := board.StartGame(3)
	recipientId := players[0].Id
	senderId := players[1].Id

	amount := 1_000

	_, err := AddTransaction(financeDomain.NewTransaction(recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}

	actualSenderId := players[2].Id
	ResolveLatestTransaction(actualSenderId)

	assert.Equal(t, 1_000, getAccountByPlayerId(senderId).Balance)
	assert.Equal(t, 0, getAccountByPlayerId(actualSenderId).Balance)
	assert.Equal(t, 2_000, getAccountByPlayerId(recipientId).Balance)
}

func TestTransactionCanOnlyBeResolvedOnce(t *testing.T) {

	InitAccounts()
	config.Init()

	players, _ := board.StartGame(3)
	recipientId := players[0].Id
	senderId := players[1].Id

	amount := 100

	_, err := AddTransaction(financeDomain.NewTransaction(recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}

	ResolveLatestTransaction(senderId)
	ResolveLatestTransaction(senderId)

	assert.Equal(t, 900, getAccountByPlayerId(senderId).Balance)
	assert.Equal(t, 1100, getAccountByPlayerId(recipientId).Balance)
}
