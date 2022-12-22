package finance

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/communication"
	financeDomain "iot-monopoly/finance/domain"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
	propertyDomain "iot-monopoly/property/domain"
)

func StartEventListeners() {

	startLapFinishedEventHandler()
	startPlayerOnOwnedFieldEventHandler()
	startPropertyTransferCreatedEventHandler()
	startGameStartedEventHandler()
	startCardWithPayoutDrewEventHandler()
	startCardWithFeeEventHandler()
}

func startLapFinishedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(*boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			addToAccount(getAccountByPlayerId(lapFinishedEvent.PlayerId).Id, 100)
		},
		Matcher: string(communication.LAP_FINISHED),
	})
}
func startGameStartedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			InitAccounts()
		},
		Matcher: string(communication.GAME_STARTED),
	})
}
func startPlayerOnOwnedFieldEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*propertyDomain.PlayerOnOwnedFieldEvent)
			AddTransaction(financeDomain.NewTransactionWithId(uuid.NewString(), transactionInformation.OwnerId, transactionInformation.PlayerId, transactionInformation.Fee))
		},
		Matcher: string(communication.PLAYER_ON_OWNED_FIELD),
	})
}

func startPropertyTransferCreatedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*propertyDomain.PropertyTransferCreatedEvent)
			AddTransaction(financeDomain.NewTransactionWithId(transactionInformation.TransactionId, transactionInformation.ReceiverId, transactionInformation.SenderId, transactionInformation.Price))
		},
		Matcher: string(communication.PROPERTY_TRANSFER_CREATED),
	})
}

func startCardWithPayoutDrewEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			payoutInformation := e.Data.(*gameEventsDomain.CardWithPayoutEvent)
			addToAccount(getAccountByPlayerId(payoutInformation.PlayerId).Id, payoutInformation.Amount)
		},
		Matcher: string(communication.CARD_WITH_PAYOUT_DREW),
	})

}
func startCardWithFeeEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*gameEventsDomain.CardWithFeeEvent)
			AddTransaction(financeDomain.NewTransactionWithId(uuid.NewString(), transactionInformation.RecipientId, transactionInformation.PlayerId, transactionInformation.Fee))
		},
		Matcher: string(communication.CARD_WITH_FEE_DREW),
	})

}
