package finance

import (
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/eventing/config"
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

func TestTransactionCanOnlyBeResolvedOnce(t *testing.T) {
	InitAccounts()
	config.Init()

	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	amount := 1_000

	transaction, err := AddTransaction(financeDomain.NewTransaction(recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}
	ResolveLatestTransaction(senderId)

	assert.Panics(t, func() {
		ResolveLatestTransaction(transaction.Id)
	})
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
	transaction := ResolveLatestTransaction(actualSenderId)

	assert.Equal(t, 1_000, getAccountByPlayerId(senderId).Balance)
	assert.Equal(t, 0, getAccountByPlayerId(actualSenderId).Balance)
	assert.Equal(t, 2_000, getAccountByPlayerId(recipientId).Balance)
	assert.Equal(t, false, transaction.IsPending())
	assert.False(t, transaction.ExecutionTime.IsZero())

}
