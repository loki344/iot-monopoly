package finance

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	financeDomain "iot-monopoly/finance/domain"
)

func StartEventListeners() {

	startLapFinishedEventHandler()
	startPlayerOnOwnedFieldEventHandler()
	startPropertyTransferCreatedEventHandler()
	startGameStartedEventHandler()
	startCreditAddedEventHandler()
}

func startLapFinishedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(*boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			addToAccount(getAccountByPlayerId(lapFinishedEvent.PlayerId).Id, 100)
		},
		Matcher: string(eventing.LAP_FINISHED),
	})
}
func startGameStartedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			InitAccounts()
		},
		Matcher: string(eventing.GAME_STARTED),
	})
}
func startPlayerOnOwnedFieldEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*boardDomain.PlayerOnOwnedFieldEvent)
			AddTransaction(financeDomain.NewTransactionWithId(uuid.NewString(), transactionInformation.OwnerId, transactionInformation.PlayerId, transactionInformation.Fee))
		},
		Matcher: string(eventing.PLAYER_ON_OWNED_FIELD),
	})
}

func startPropertyTransferCreatedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*boardDomain.PropertyTransferCreatedEvent)
			AddTransaction(financeDomain.NewTransactionWithId(transactionInformation.TransactionId, transactionInformation.ReceiverId, transactionInformation.SenderId, transactionInformation.Price))
		},
		Matcher: string(eventing.PROPERTY_TRANSFER_CREATED),
	})
}

func startCreditAddedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			creditAddedEvent := e.Data.(*boardDomain.CreditAddedEvent)
			addToAccount(creditAddedEvent.RecipientAccountId, creditAddedEvent.Amount)
		},
		Matcher: string(eventing.PAYOUT_REQUESTED),
	})

}
