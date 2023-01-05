package adapter

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	"iot-monopoly/game/domain"
	"iot-monopoly/game/domain/events"
)

func StartEventListeners() {

	startTransactionResolvedHandler()
	startGameEventWithPayoutAcceptedHandler()
	startGameEventWithFeeAcceptedHandler()
}

func startTransactionResolvedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionResolved := e.Data.(events.TransactionResolvedEvent)
			fmt.Printf("Transaction %s is resolved, check for pending actions to trigger\n", transactionResolved.TransactionId)
			GetCurrentGame().TransferOwnership(transactionResolved.TransactionId)
		},
		Matcher: string(eventing.TRANSACTION_RESOLVED),
	})
}

func startGameEventWithPayoutAcceptedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			payoutInformation := e.Data.(*events.GameEventWithPayoutAcceptedEvent)
			GetCurrentGame().FindAccountById(payoutInformation.RecipientId).Deposit(payoutInformation.Amount)
		},
		Matcher: string(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED),
	})

}
func startGameEventWithFeeAcceptedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*events.GameEventWithFeeAcceptedEvent)
			domain.NewTransaction(uuid.NewString(), transactionInformation.RecipientId, transactionInformation.PayerId, transactionInformation.Fee)
		},
		Matcher: string(eventing.GAME_EVENT_WITH_FEE_ACCEPTED),
	})
}
