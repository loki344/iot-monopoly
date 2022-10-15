package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/domain"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	StartEventHandler()

	id := board.StartGame(1)[0].Id
	eventing.FireEvent(eventing.LAP_FINISHED, eventingDomain.LapFinishedEvent{PlayerId: id})
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

	eventing.FireEvent(eventing.TRANSACTION_REQUESTED, eventingDomain.NewTransactionRequest(transactionId, recipientId, senderId, amount))

	transaction := GetTransaction(transactionId)
	assert.Equal(t, 1, len(getPendingTransactions()))
	assert.Equal(t, transactionId, transaction.Id())
	assert.Equal(t, true, transaction.IsPending())
	assert.Equal(t, senderId, transaction.SenderId())
	assert.Equal(t, recipientId, transaction.RecipientId())
	assert.Equal(t, amount, transaction.Amount())
}
