package main

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
	"iot-monopoly/finance"
	"iot-monopoly/movement"
)

func StartGame(players []boardDomain.Player) {
	eventing.StartExternalEventbus()
	eventing.StartInternalEventBus()
	movement.StartEventbus()
	finance.StartEventbus()
	board.InitBoard(nil, players)
}

func EndGame() {
	eventing.CloseExternalEventbus()
	eventing.CloseInternalEventbus()
	movement.CloseEventbus()
	finance.CloseEventbus()
	board.ClearBoard()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {
	//TODO get player infos from board
	StartGame([]boardDomain.Player{{1, 0, 1000}, {2, 0, 1000}})

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
	PrintGameState()
}
func PrintGameState() {

	fmt.Println("-------------------------------------")
	fmt.Println("Game State")
	fmt.Println("Players:")
	fmt.Println(board.Players)
	fmt.Println("Fields:")
	fmt.Println(board.Fields)
	fmt.Println("-------------------------------------")
}
