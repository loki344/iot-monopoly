package movement

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
)

func StartEventbus() {

	tr := bus.GetBus()
	channel := "external"

	sensorEventHandler, err := tr.ListenRequestStream(channel)
	if err != nil {
		fmt.Println(err)
	}
	sensorEventHandler.Handle(
		func(msg *model.Message) {
			sensorEvent := msg.Payload.(eventing.SensorEvent)
			board.GetPlayer(sensorEvent.PlayerId).MovePlayer(sensorEvent.FieldIndex)
		},
		func(err error) {
			fmt.Println(err)
		})
}
