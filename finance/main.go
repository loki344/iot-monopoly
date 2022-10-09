package finance

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/game_master"
	"iot-monopoly/movement"
)

func Init() {

	tr := bus.GetBus()
	channel := "lapFinished"

	sensorEventHandler, err := tr.ListenRequestStream(channel)
	if err != nil {
		fmt.Println(err)
	}
	sensorEventHandler.Handle(
		func(msg *model.Message) {
			lapFinishedEvent := msg.Payload.(movement.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			game_master.GetPlayer(lapFinishedEvent.PlayerId).Balance += 100
		},
		func(err error) {
			fmt.Println(err)
		})
}
