package finance

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	StartEventHandler()

	id := uuid.New().String()
	board.StartGame([]boardDomain.Player{{id, 0, 1000}})
	eventing.FireEvent(eventing.LAP_FINISHED, eventing.LapFinishedEvent{PlayerId: id})
	assert.Equal(t, 1100, board.GetPlayer(id).Balance)
}

func TestTransactionRequestedAppearsInPendingTransactions(t *testing.T) {

	StartEventHandler()

	recipientId := uuid.New().String()
	senderId := uuid.New().String()
	amount := 1000

	transactionRequest := eventing.NewTransactionRequest(recipientId, senderId, amount)
	eventing.FireEvent(eventing.TRANSACTION_REQUESTED, transactionRequest)
	assert.Equal(t, 1, len(getPendingTransactions()))

	transaction := GetTransaction(transactionRequest.Id)
	assert.Equal(t, recipientId, transaction.RecipientId())
	assert.Equal(t, senderId, transaction.SenderId())
	assert.Equal(t, amount, transaction.Amount())
	assert.Equal(t, true, transaction.IsPending())
}
