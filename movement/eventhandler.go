package movement

import (
	"fmt"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
)

func StartEventHandler() {

	channel := eventing.EXTERNAL

	sensorEventHandler := eventing.ListenRequestStream(channel)

	sensorEventHandler.Handle(
		func(msg *model.Message) {
			sensorEvent := msg.Payload.(eventing.SensorEvent)
			board.MovePlayer(sensorEvent.PlayerId, sensorEvent.FieldIndex)
		},
		func(err error) {
			fmt.Println(err)
		})
}
