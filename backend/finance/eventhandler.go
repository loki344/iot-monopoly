package finance

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	financeDomain "iot-monopoly/finance/domain"
)

func StartEventListeners() {

	startLapFinishedEventHandler()
	startTransactionRequestedEventHandler()
}

func startLapFinishedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			board.GetPlayer(lapFinishedEvent.PlayerId).Balance += 100
		},
		Matcher: string(eventing.LAP_FINISHED),
	})
}

func startTransactionRequestedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionRequest := e.Data.(financeDomain.TransactionRequest)
			_, err := AddTransaction(financeDomain.NewTransaction(transactionRequest.RecipientId, transactionRequest.SenderId, transactionRequest.Amount))
			if err != nil {
				fmt.Println(err)
				return
			}
		},
		Matcher: string(eventing.TRANSACTION_REQUESTED),
	})
}
