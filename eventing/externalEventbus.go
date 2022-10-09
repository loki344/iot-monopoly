package eventing

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
)

func StartExternalEventbus() {

	// TODO connect to rabbitMQ

	//TODO remove as soon as we're connected to real eventbus
	initExternalEventbus()
}
func initExternalEventbus() {

	tr := bus.GetBus()
	channel := "external"

	sensorEventHandler, err := tr.ListenRequestStream(channel)
	if err != nil {
		fmt.Println(err)
	}
	sensorEventHandler.Handle(
		func(msg *model.Message) {
			fmt.Println("received message on EXTERNAL channel")
			sensorEvent := msg.Payload.(SensorEvent)
			fmt.Println(sensorEvent)
		},
		func(err error) {
			fmt.Println(err)
		})
}

type SensorEvent struct {
	PlayerId   int
	FieldIndex int
}
