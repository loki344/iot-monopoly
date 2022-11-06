package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
	"iot-monopoly/finance/domain"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	StartEventHandler()

	id := board.StartGame(1)[0].Id
	eventing.FireEvent(eventing.LAP_FINISHED, domain.LapFinishedEvent{PlayerId: id})
	assert.Equal(t, 1100, board.GetPlayer(id).Balance)
}

//TODO transactionrequest turns into transaction

func TestTransactionRequestTurnsIntoTransaction(t *testing.T) {

	startTransactionRequestedEventHandler()
	players := board.StartGame(2)
	recipientId := players[0].Id
	senderId := players[1].Id

	transactionId := uuid.NewString()
	amount := 1_000

	eventing.FireEvent(eventing.TRANSACTION_REQUESTED, domain.NewTransactionRequest(transactionId, recipientId, senderId, amount))

	var transaction financeDomain.Transaction
	assert.NotPanics(t, func() {
		transaction = GetTransaction(transactionId)
	})
	assert.Equal(t, transactionId, transaction.Id())
	assert.Equal(t, true, transaction.IsPending())
	assert.Equal(t, senderId, transaction.SenderId())
	assert.Equal(t, recipientId, transaction.RecipientId())
	assert.Equal(t, amount, transaction.Amount())
}
