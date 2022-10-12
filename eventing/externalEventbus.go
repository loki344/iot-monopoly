package eventing

import (
	"fmt"
	"github.com/vmware/transport-go/model"
)

func StartExternalEventHandler() {

	// TODO connect to rabbitMQ

	//TODO remove as soon as we're connected to real eventbus
	initExternalEventbus()
}
func initExternalEventbus() {

	channel := EXTERNAL

	sensorEventHandler := ListenRequestStream(channel)

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
	PlayerId   string
	FieldIndex int
}
