package adapter

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	"iot-monopoly/finance/adapter/repository"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
	boardDomain "iot-monopoly/player/domain"
	propertyDomain "iot-monopoly/property/domain"
)

func StartEventListeners() {

	startLapFinishedEventHandler()
	startPlayerOnOwnedFieldEventHandler()
	startPropertyTransferCreatedEventHandler()
	startGameStartedEventHandler()
	startGameEventWithPayoutAcceptedHandler()
	startGameEventWithFeeAcceptedHandler()
}

func startLapFinishedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(*boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			repository.GetAccountByPlayerId(lapFinishedEvent.PlayerId).Add(100)
		},
		Matcher: string(eventing.LAP_FINISHED),
	})
}
func startGameStartedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			repository.InitAccounts()
		},
		Matcher: string(eventing.GAME_STARTED),
	})
}
func startPlayerOnOwnedFieldEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*propertyDomain.PlayerOnOwnedFieldEvent)
			AddTransaction(uuid.NewString(), transactionInformation.OwnerId, transactionInformation.PlayerId, transactionInformation.Fee)
		},
		Matcher: string(eventing.PLAYER_ON_OWNED_FIELD),
	})
}

func startPropertyTransferCreatedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*propertyDomain.PropertyTransferCreatedEvent)
			AddTransaction(transactionInformation.TransactionId, transactionInformation.ReceiverId, transactionInformation.SenderId, transactionInformation.Price)
		},
		Matcher: string(eventing.PROPERTY_TRANSFER_CREATED),
	})
}

func startGameEventWithPayoutAcceptedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			payoutInformation := e.Data.(*gameEventsDomain.GameEventWithPayoutAcceptedEvent)
			repository.GetAccountByPlayerId(payoutInformation.PlayerId).Add(payoutInformation.Amount)
		},
		Matcher: string(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED),
	})

}
func startGameEventWithFeeAcceptedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*gameEventsDomain.GameEventWithFeeAcceptedEvent)
			AddTransaction(uuid.NewString(), transactionInformation.RecipientId, transactionInformation.PlayerId, transactionInformation.Fee)
		},
		Matcher: string(eventing.GAME_EVENT_WITH_FEE_ACCEPTED),
	})

}
