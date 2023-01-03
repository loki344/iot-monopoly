package adapter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/eventing/config"
	"iot-monopoly/finance/adapter/repository"
	"testing"
)

func TestTransactionWithInsufficientBalance(t *testing.T) {
	config.Init()
	repository.InitAccounts()

	recipientId := GetAccounts()[0].PlayerId
	senderId := GetAccounts()[1].PlayerId

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	err := ResolveLatestTransaction(GetAccounts()[1].Id)
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidTransaction(t *testing.T) {
	config.Init()
	repository.InitAccounts()

	recipientId := GetAccounts()[0].PlayerId
	senderId := GetAccounts()[1].PlayerId

	amount := 1_000

	transaction := AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	assert.Equal(t, recipientId, transaction.RecipientId)
	assert.Equal(t, senderId, transaction.SenderId)
	assert.Equal(t, amount, transaction.Amount)
}

func TestResolveTransactionChangesBalance(t *testing.T) {
	config.Init()
	repository.InitAccounts()

	recipientId := GetAccounts()[0].PlayerId
	senderId := GetAccounts()[1].PlayerId

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	ResolveLatestTransaction(GetAccounts()[1].Id)

	assert.Equal(t, 0, repository.GetAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 2_000, repository.GetAccountByPlayerId(recipientId).Balance())
}

func TestTransactionWithChangedSenderId(t *testing.T) {

	config.Init()
	repository.InitAccounts()

	recipientId := GetAccounts()[0].PlayerId
	senderId := GetAccounts()[1].PlayerId

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	actualSenderId := GetAccounts()[2].Id
	ResolveLatestTransaction(actualSenderId)

	assert.Equal(t, 1_000, repository.GetAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 0, repository.GetAccountByPlayerId(GetAccounts()[2].PlayerId).Balance())
	assert.Equal(t, 2_000, repository.GetAccountByPlayerId(recipientId).Balance())
}

func TestTransactionCanOnlyBeResolvedOnce(t *testing.T) {

	repository.InitAccounts()
	config.Init()

	recipientId := GetAccounts()[0].PlayerId
	senderId := GetAccounts()[1].PlayerId

	amount := 100

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	ResolveLatestTransaction(GetAccounts()[1].Id)
	ResolveLatestTransaction(GetAccounts()[1].Id)

	assert.Equal(t, 900, repository.GetAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 1100, repository.GetAccountByPlayerId(recipientId).Balance())
}
