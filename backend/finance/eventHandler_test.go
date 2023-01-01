package finance

import (
	"context"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/communication"
	"iot-monopoly/communication/config"
	financeDomain "iot-monopoly/finance/domain"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
	"iot-monopoly/player"
	boardDomain "iot-monopoly/player/domain"
	"iot-monopoly/property/domain"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	config.Init()
	StartEventListeners()

	players, _ := player.Init(1)
	id := players[0].Id()
	communication.FireEvent(communication.LAP_FINISHED, &boardDomain.LapFinishedEvent{PlayerId: id})
	assert.Equal(t, 1100, getAccountByPlayerId(id).Balance())
}

func TestPlayerOnOwnedFieldFiresTransactionRequestEvent(t *testing.T) {

	config.Init()
	StartEventListeners()
	payerId := "Player_1"
	ownerId := "Player_2"

	var receivedEvents = 0
	const price = 1000
	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transaction := e.Data.(financeDomain.TransactionCreatedEvent)
			assert.Equal(t, payerId, transaction.Transaction.SenderId())
			assert.Equal(t, ownerId, transaction.Transaction.RecipientId())
			assert.Equal(t, price, transaction.Transaction.Amount())
			receivedEvents++
		},
		Matcher: string(communication.TRANSACTION_CREATED),
	})

	var tempFinancialDetails = &propertyDomain.FinancialDetails{100, 100, 100, propertyDomain.Revenue{1000, 200, 300, 400, 500, 800}}
	property := propertyDomain.NewPropertyField("Property green 2", uuid.NewString(), tempFinancialDetails)
	property.SetOwnerId(ownerId)

	property.OnPlayerEnter(payerId)

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerReceivesMoneyWhenCardWithPayoutDrewEventFired(t *testing.T) {

	config.Init()
	StartEventListeners()

	players, _ := player.Init(1)
	initAccounts()
	currentPlayer := players[0]
	communication.FireEvent(communication.CARD_WITH_PAYOUT_ACCEPTED, gameEventsDomain.NewCardWithPayoutDrewEvent(currentPlayer.Id(), 200))
	assert.Equal(t, 1200, getAccountByPlayerId(currentPlayer.Id()).Balance())
}
