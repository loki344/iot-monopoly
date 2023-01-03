package adapter

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	financeDomain "iot-monopoly/finance/domain"
	boardDomain "iot-monopoly/player/domain"
	"iot-monopoly/property/adapter/repository"
	"strconv"
)

func StartEventListeners() {

	startTransactionResolvedHandler()
	startGameStartedEventHandler()
	startPlayerMovedEventHandler()
}

func startGameStartedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			repository.InitFields()
		},
		Matcher: string(eventing.GAME_STARTED),
	})
}

func startTransactionResolvedHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionResolved := e.Data.(financeDomain.TransactionResolvedEvent)
			fmt.Printf("Transaction %s is resolved, check for pending actions to trigger\n", transactionResolved.TransactionId)
			transferOwnerShip(transactionResolved.TransactionId)
		},
		Matcher: string(eventing.TRANSACTION_RESOLVED),
	})
}

func startPlayerMovedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			playerMovedEvent := e.Data.(*boardDomain.PlayerMovedEvent)

			property := getPropertyById(strconv.FormatInt(int64(playerMovedEvent.FieldIndex), 10))

			if property != nil {
				property.OnPlayerEnter(playerMovedEvent.PlayerId)

			}
		},
		Matcher: string(eventing.PLAYER_MOVED),
	})
}
