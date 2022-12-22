package property

import (
	"context"
	"errors"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/communication"
	financeDomain "iot-monopoly/finance/domain"
	"strconv"
)

func StartEventListeners() {

	startTransactionResolvedHandler()
	startGameStartedEventHandler()
	startPlayerMovedEventHandler()
}

func startGameStartedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			initFields()
		},
		Matcher: string(communication.GAME_STARTED),
	})
}

func startTransactionResolvedHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionResolved := e.Data.(financeDomain.TransactionResolvedEvent)
			fmt.Printf("Transaction %s is resolved, check for pending actions to trigger\n", transactionResolved.TransactionId)
			transferOwnerShip(transactionResolved.TransactionId)
		},
		Matcher: string(communication.TRANSACTION_RESOLVED),
	})
}

func startPlayerMovedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			playerMovedEvent := e.Data.(boardDomain.PlayerMovedEvent)

			totalFieldCount := GetFieldsCount()
			if playerMovedEvent.FieldIndex > totalFieldCount-1 || playerMovedEvent.FieldIndex < 0 {
				errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", playerMovedEvent.FieldIndex, totalFieldCount))
			}

			GetFieldById(strconv.FormatInt(int64(playerMovedEvent.FieldIndex), 10)).OnPlayerEnter(playerMovedEvent.PlayerId)
		},
		Matcher: string(communication.TRANSACTION_RESOLVED),
	})
}
