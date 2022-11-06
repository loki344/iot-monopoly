package finance

import (
	"fmt"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"iot-monopoly/finance/domain"
)

var started = false

func StartEventHandler() {

	if !started {
		startLapFinishedEventHandler()
		startTransactionRequestedEventHandler()
		started = true
	}
}

func startLapFinishedEventHandler() {
	channel := eventing.LAP_FINISHED

	lapFinishedEventHandler := eventing.ListenRequestStream(channel)

	lapFinishedEventHandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			board.GetPlayer(lapFinishedEvent.PlayerId).Balance += 100
		},
		func(err error) {
			fmt.Println(err)
		})
}

func startTransactionRequestedEventHandler() {
	channel := eventing.TRANSACTION_REQUESTED

	transactionRequestedHandler := eventing.ListenRequestStream(channel)

	transactionRequestedHandler.Handle(
		func(msg *model.Message) {
			transactionRequest := msg.Payload.(financeDomain.TransactionRequested)
			//TODO handle error
			AddTransaction(*financeDomain.NewTransaction(transactionRequest.Id(), transactionRequest.RecipientId(), transactionRequest.SenderId(), transactionRequest.Amount()))
		},
		func(err error) {
			fmt.Println(err)
		})
}
