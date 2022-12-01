package finance

import (
	"context"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/config"
	financeDomain "iot-monopoly/finance/domain"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	config.Init()
	StartEventListeners()

	players, _ := board.StartGame(1)
	id := players[0].Id
	eventing.FireEvent(eventing.LAP_FINISHED, &boardDomain.LapFinishedEvent{PlayerId: id})
	assert.Equal(t, 1100, getAccountByPlayerId(id).Balance)
}

func TestPlayerOnOwnedFieldFiresTransactionRequestEvent(t *testing.T) {

	config.Init()
	StartEventListeners()
	payerId := "Player_1"
	ownerId := "Player_2"

	var receivedEvents = 0
	const price = 1000
	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transaction := e.Data.(financeDomain.TransactionAddedEvent)
			assert.Equal(t, payerId, transaction.Transaction.SenderId)
			assert.Equal(t, ownerId, transaction.Transaction.RecipientId)
			assert.Equal(t, price, transaction.Transaction.Amount)
			receivedEvents++
		},
		Matcher: string(eventing.TRANSACTION_REQUEST),
	})

	var tempFinancialDetails = &boardDomain.FinancialDetails{100, 100, 100, boardDomain.Revenue{1000, 200, 300, 400, 500, 800}}
	property := boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{"Property green 2", uuid.NewString()}, tempFinancialDetails)
	property.OwnerId = ownerId

	property.OnPlayerEnter(&boardDomain.Player{payerId, 0, "Account_Player_1"})

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerReceivesMoneyWhenPaymentEventFired(t *testing.T) {

	config.Init()
	StartEventListeners()

	players, _ := board.StartGame(1)
	player := players[0]
	eventing.FireEvent(eventing.PAYOUT_REQUESTED, boardDomain.NewCreditAddedEvent(player.AccountId, 200))
	assert.Equal(t, 1200, getAccountByPlayerId(player.Id).Balance)
}
