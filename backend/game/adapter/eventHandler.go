package adapter

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	"iot-monopoly/game/domain"
)

func StartEventListeners() {

	startTransactionResolvedHandler()
}

func startTransactionResolvedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionResolved := e.Data.(domain.TransactionResolvedEvent)
			fmt.Printf("Transaction %s is resolved, check for pending actions to trigger\n", transactionResolved.TransactionId)
			GetCurrentGame().TransferOwnership(transactionResolved.TransactionId)
		},
		Matcher: string(eventing.TRANSACTION_RESOLVED),
	})
}

// TODO we need a eventhandler for creating transactions from gameEvents
