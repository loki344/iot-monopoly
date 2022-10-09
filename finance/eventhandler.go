package finance

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
)

func StartEventbus() {

	tr := bus.GetBus()
	channel := "lapFinished"

	sensorEventHandler, err := tr.ListenRequestStream(channel)
	if err != nil {
		fmt.Println(err)
	}
	sensorEventHandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			board.GetPlayer(lapFinishedEvent.PlayerId).Balance += 100
		},
		func(err error) {
			fmt.Println(err)
		})
}
