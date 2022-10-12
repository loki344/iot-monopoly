package finance

import (
	"fmt"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
)

func StartEventHandler() {

	channel := eventing.LAP_FINISHED

	lapFinishedEventhandler := eventing.ListenRequestStream(channel)

	lapFinishedEventhandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			board.GetPlayer(lapFinishedEvent.PlayerId).Balance += 100
		},
		func(err error) {
			fmt.Println(err)
		})
}
