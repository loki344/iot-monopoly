package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/finance/financeDomain"
	"testing"
)

func TestTransactionWithInsufficientBalance(t *testing.T) {
	recipientId := uuid.NewString()
	senderId := uuid.NewString()
	board.StartGame([]boardDomain.Player{{recipientId, 0, 1000}, {senderId, 0, 500}})

	transactionId := uuid.NewString()
	amount := 1_000

	err := AddTransaction(*financeDomain.NewTransaction(transactionId, recipientId, senderId, amount))
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidTransaction(t *testing.T) {

	recipientId := uuid.NewString()
	senderId := uuid.NewString()
	board.StartGame([]boardDomain.Player{{recipientId, 0, 1000}, {senderId, 0, 1000}})

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

	recipientId := uuid.NewString()
	senderId := uuid.NewString()
	const balance = 1_000
	board.StartGame([]boardDomain.Player{{recipientId, 0, balance}, {senderId, 0, balance}})

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
	recipientId := uuid.NewString()
	senderId := uuid.NewString()
	const balance = 1_000
	board.StartGame([]boardDomain.Player{{recipientId, 0, balance}, {senderId, 0, balance}})

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
