package adapter

import (
	"context"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"github.com/stretchr/testify/assert"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/config"
	"iot-monopoly/finance/adapter/repository"
	financeDomain "iot-monopoly/finance/domain"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
	playerDomain "iot-monopoly/player/domain"
	"iot-monopoly/property/domain"
	"testing"
)

func TestPlayerReceivesMoneyWhenLapFinished(t *testing.T) {

	config.Init()
	StartEventListeners()
	repository.InitAccounts()

	eventing.FireEvent(eventing.LAP_FINISHED, &playerDomain.LapFinishedEvent{PlayerId: GetAccounts()[0].PlayerId})
	assert.Equal(t, 1100, repository.GetAccountByPlayerId(GetAccounts()[0].PlayerId).Balance())
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
			transaction := e.Data.(financeDomain.TransactionCreatedEvent)
			assert.Equal(t, payerId, transaction.Transaction.SenderId)
			assert.Equal(t, ownerId, transaction.Transaction.RecipientId)
			assert.Equal(t, price, transaction.Transaction.Amount)
			receivedEvents++
		},
		Matcher: string(eventing.TRANSACTION_CREATED),
	})

	var tempFinancialDetails = &domain.FinancialDetails{100, 100, 100, domain.Revenue{1000, 200, 300, 400, 500, 800}}
	property := domain.NewPropertyField("Property green 2", uuid.NewString(), tempFinancialDetails)
	property.OwnerId = ownerId

	property.OnPlayerEnter(payerId)

	assert.Equal(t, 1, receivedEvents)
}

func TestPlayerReceivesMoneyWhenCardWithPayoutDrewEventFired(t *testing.T) {

	config.Init()
	StartEventListeners()

	repository.InitAccounts()
	eventing.FireEvent(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED, gameEventsDomain.NewGameEventWithPayoutAcceptedEvent(GetAccounts()[0].PlayerId, 200))
	assert.Equal(t, 1200, repository.GetAccountByPlayerId(GetAccounts()[0].PlayerId).Balance())
}
