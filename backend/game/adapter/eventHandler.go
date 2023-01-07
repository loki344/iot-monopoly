package adapter

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	"iot-monopoly/game/adapter/repository"
	"iot-monopoly/game/domain"
)

func StartEventListeners() {

	startTransactionResolvedHandler()
	startGameEventWithPayoutAcceptedHandler()
	startGameEventWithFeeAcceptedHandler()
	startPlayerOnOwnedFieldEvent()
}

func startTransactionResolvedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionResolved := e.Data.(domain.TransactionResolvedEvent)
			fmt.Printf("Transaction %s is resolved, check for pending actions to trigger\n", transactionResolved.TransactionId)
			GetCurrentGame().ResolvePendingPropertyTransfer(transactionResolved.TransactionId)
		},
		Matcher: string(eventing.TRANSACTION_RESOLVED),
	})
}

func startGameEventWithPayoutAcceptedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			payoutInformation := e.Data.(*domain.GameEventWithPayoutAcceptedEvent)
			repository.FindAccountById(payoutInformation.RecipientId).Deposit(payoutInformation.Amount)
		},
		Matcher: string(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED),
	})

}
func startGameEventWithFeeAcceptedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*domain.GameEventWithFeeAcceptedEvent)
			repository.CreateTransaction(domain.NewTransaction(uuid.NewString(), transactionInformation.RecipientId, transactionInformation.PayerId, transactionInformation.Fee))
		},
		Matcher: string(eventing.GAME_EVENT_WITH_FEE_ACCEPTED),
	})
}

func startPlayerOnOwnedFieldEvent() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionInformation := e.Data.(*domain.PlayerOnOwnedFieldEvent)
			repository.CreateTransaction(domain.NewTransaction(uuid.NewString(), transactionInformation.OwnerId, transactionInformation.PlayerId, transactionInformation.Fee))
		},
		Matcher: string(eventing.PLAYER_ON_OWNED_FIELD),
	})
}
