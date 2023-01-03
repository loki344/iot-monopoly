package financeAdapter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/communication/config"
	"iot-monopoly/finance/adapter/repository"
	adapter "iot-monopoly/player/adapter"
	"testing"
)

func TestTransactionWithInsufficientBalance(t *testing.T) {
	config.Init()
	repository.InitAccounts()

	players, _ := adapter.Init(2)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	err := ResolveLatestTransaction(players[1].AccountId())
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidTransaction(t *testing.T) {
	config.Init()
	repository.InitAccounts()

	players, _ := adapter.Init(2)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	transaction := AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	assert.Equal(t, recipientId, transaction.RecipientId)
	assert.Equal(t, senderId, transaction.SenderId)
	assert.Equal(t, amount, transaction.Amount)
}

func TestResolveTransactionChangesBalance(t *testing.T) {
	config.Init()
	repository.InitAccounts()

	players, _ := adapter.Init(2)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	ResolveLatestTransaction(players[1].AccountId())

	assert.Equal(t, 0, repository.GetAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 2_000, repository.GetAccountByPlayerId(recipientId).Balance())
}

func TestTransactionWithChangedSenderId(t *testing.T) {

	config.Init()
	repository.InitAccounts()

	players, _ := adapter.Init(3)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	actualSenderId := players[2].AccountId()
	ResolveLatestTransaction(actualSenderId)

	assert.Equal(t, 1_000, repository.GetAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 0, repository.GetAccountByPlayerId(players[2].Id()).Balance())
	assert.Equal(t, 2_000, repository.GetAccountByPlayerId(recipientId).Balance())
}

func TestTransactionCanOnlyBeResolvedOnce(t *testing.T) {

	repository.InitAccounts()
	config.Init()

	players, _ := adapter.Init(3)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 100

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	ResolveLatestTransaction(players[1].AccountId())
	ResolveLatestTransaction(players[1].AccountId())

	assert.Equal(t, 900, repository.GetAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 1100, repository.GetAccountByPlayerId(recipientId).Balance())
}