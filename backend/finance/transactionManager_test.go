package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/finance/domain"
	"testing"
)

func TestTransactionWithInsufficientBalance(t *testing.T) {
	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	transactionId := uuid.NewString()
	amount := 1_000

	err := AddTransaction(*financeDomain.NewTransaction(transactionId, recipientId, senderId, amount))
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidTransaction(t *testing.T) {

	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	transactionId := uuid.NewString()
	amount := 1_000

	err := AddTransaction(*financeDomain.NewTransaction(transactionId, recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}
	transaction := GetTransaction(transactionId)

	assert.Equal(t, recipientId, transaction.RecipientId())
	assert.Equal(t, senderId, transaction.SenderId())
	assert.Equal(t, amount, transaction.Amount())
	assert.Equal(t, transactionId, transaction.Id())
}

func TestResolveTransactionChangesBalance(t *testing.T) {

	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	transactionId := uuid.NewString()
	amount := 1_000

	err := AddTransaction(*financeDomain.NewTransaction(transactionId, recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}

	ResolveTransaction(transactionId)

	assert.Equal(t, 0, board.GetPlayer(senderId).Balance)
	assert.Equal(t, 2_000, board.GetPlayer(recipientId).Balance)
}

func TestTransactionCanOnlyBeResolvedOnce(t *testing.T) {
	players, _ := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	transactionId := uuid.NewString()
	amount := 1_000

	err := AddTransaction(*financeDomain.NewTransaction(transactionId, recipientId, senderId, amount))
	if err != nil {
		assert.NoError(t, err)
	}
	ResolveTransaction(transactionId)

	assert.Panics(t, func() {
		ResolveTransaction(transactionId)
	})
}
