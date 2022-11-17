package finance

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	financeDomain "iot-monopoly/finance/domain"
)

func StartEventListeners() {

	startLapFinishedEventHandler()
	startPropertyFeeEventHandler()
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
func startPropertyFeeEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			transactionRequest := e.Data.(*boardDomain.TransactionRequest)
			AddTransaction(financeDomain.NewTransactionWithId(transactionRequest.TransactionId, transactionRequest.ReceiverId, transactionRequest.SenderId, transactionRequest.Price))
		},
		Matcher: "((^|, )(propertyTransactionStarted|propertyFee))+$",
	})
}
