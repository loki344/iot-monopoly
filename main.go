package main

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"iot-monopoly/eventing"
	"iot-monopoly/game_master"
)

func Init() {
	eventing.ConnectToRabbitMQ()
	eventing.StartInternalEventBus()
	game_master.Init()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {
	Init()

	//TODO remove later
	fmt.Println("Sending external Events!!")
	events := []eventing.SensorEvent{{1, 5}}
	ts := bus.GetBus()
	for _, event := range events {
		handler, err := ts.RequestOnce("external", event)
		if err != nil {
			//TODO something went wrong
			return
		}
		handler.Fire()
	}
}
