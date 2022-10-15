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
}

var transactions []financeDomain.Transaction

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

func AddTransaction(transaction financeDomain.Transaction) {
	fmt.Printf("Adding transaction %s to pending transactions\n", transaction.Id())
	transactions = append(transactions, transaction)
	if transaction.IsPending() {
		eventing.FireEvent(eventing.TRANSACTION_REQUESTED, transaction)
	}
}

func GetTransaction(id string) *financeDomain.Transaction {

	for _, transaction := range transactions {
		if transaction.Id() == id {
			return &transaction
		}
	}

	panic(fmt.Sprintf("no transaction found with id %s", id))
}

func getPendingTransactions() []financeDomain.Transaction {
	return transactions
}
