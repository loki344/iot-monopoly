package board

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
)

func StartEventListeners() {

	startTransactionResolvedHandler()
}

func startTransactionResolvedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionId := e.Data.(string)
			fmt.Printf("Transaction %s is resolved, check for pending actions to trigger", transactionId)
			transferOwnerShip(transactionId)
		},
		Matcher: string(eventing.TRANSACTION_RESOLVED),
	})
}
