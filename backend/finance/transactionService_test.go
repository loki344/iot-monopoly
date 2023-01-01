package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/communication/config"
	"iot-monopoly/player"
	"testing"
)

func TestTransactionWithInsufficientBalance(t *testing.T) {
	config.Init()
	initAccounts()
	players, _ := player.Init(2)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	err := ResolveLatestTransaction(senderId)
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidTransaction(t *testing.T) {
	config.Init()
	initAccounts()

	players, _ := player.Init(2)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	transaction := AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	assert.Equal(t, recipientId, transaction.RecipientId())
	assert.Equal(t, senderId, transaction.SenderId())
	assert.Equal(t, amount, transaction.Amount())
}

func TestResolveTransactionChangesBalance(t *testing.T) {
	config.Init()
	initAccounts()

	players, _ := player.Init(2)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	ResolveLatestTransaction(senderId)

	assert.Equal(t, 0, getAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 2_000, getAccountByPlayerId(recipientId).Balance())
}

func TestTransactionWithChangedSenderId(t *testing.T) {

	initAccounts()
	config.Init()

	players, _ := player.Init(3)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 1_000

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	actualSenderId := players[2].Id()
	ResolveLatestTransaction(actualSenderId)

	assert.Equal(t, 1_000, getAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 0, getAccountByPlayerId(actualSenderId).Balance())
	assert.Equal(t, 2_000, getAccountByPlayerId(recipientId).Balance())
}

func TestTransactionCanOnlyBeResolvedOnce(t *testing.T) {

	initAccounts()
	config.Init()

	players, _ := player.Init(3)
	recipientId := players[0].Id()
	senderId := players[1].Id()

	amount := 100

	AddTransaction(uuid.NewString(), recipientId, senderId, amount)

	ResolveLatestTransaction(senderId)
	ResolveLatestTransaction(senderId)

	assert.Equal(t, 900, getAccountByPlayerId(senderId).Balance())
	assert.Equal(t, 1100, getAccountByPlayerId(recipientId).Balance())
}
