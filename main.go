package main

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"iot-monopoly/eventing"
	"iot-monopoly/finance"
	"iot-monopoly/game_master"
	"iot-monopoly/movement"
)

func Init() {
	eventing.ConnectToRabbitMQ()
	eventing.StartInternalEventBus()
	movement.Init()
	finance.Init()
	game_master.Init()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {
	Init()

	//TODO remove later
	fmt.Println("Sending external Events!!")
	events := []eventing.SensorEvent{
		{1, 5},
		{2, 3},
		{1, 9},
		{2, 7},
		{1, 15},
		{2, 12},
		{1, 3},
		{2, 1},
	}
	ts := bus.GetBus()
	for _, event := range events {
		handler, err := ts.RequestOnce("external", event)
		if err != nil {
			//TODO something went wrong
			return
		}
		handler.Fire()
	}
	printGameState()
}
func printGameState() {

	fmt.Println("-------------------------------------")
	fmt.Println("Game State")
	fmt.Println("Players:")
	fmt.Println(game_master.Players)
	fmt.Println("Fields:")
	fmt.Println(game_master.Fields)
	fmt.Println("-------------------------------------")
}
