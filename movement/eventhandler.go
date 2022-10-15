package movement

import (
	"fmt"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/board"
	"iot-monopoly/eventing"
	eventingDomain "iot-monopoly/eventing/domain"
)

func StartEventHandler() {

	channel := eventing.MOVEMENT

	sensorEventHandler := eventing.ListenRequestStream(channel)

	sensorEventHandler.Handle(
		func(msg *model.Message) {
			sensorEvent := msg.Payload.(eventingDomain.SensorEvent)
			board.MovePlayer(sensorEvent.PlayerId, sensorEvent.FieldIndex)
		},
		func(err error) {
			fmt.Println(err)
		})
}
