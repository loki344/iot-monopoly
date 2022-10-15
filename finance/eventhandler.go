package finance

import (
	"fmt"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/domain"
	"iot-monopoly/finance/financeDomain"
)

func StartEventHandler() {

	startLapFinishedEventHandler()
	startTransactionRequestedEventHandler()
}

func startLapFinishedEventHandler() {
	channel := eventing.LAP_FINISHED

	lapFinishedEventHandler := eventing.ListenRequestStream(channel)

	lapFinishedEventHandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(eventingDomain.LapFinishedEvent)
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
			transactionRequest := msg.Payload.(eventingDomain.TransactionRequested)
			//TODO handle error
			AddTransaction(*financeDomain.NewTransaction(transactionRequest.Id, transactionRequest.RecipientId, transactionRequest.SenderId, transactionRequest.Amount))
		},
		func(err error) {
			fmt.Println(err)
		})
}
