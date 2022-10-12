package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vmware/transport-go/bus"
	"iot-monopoly/board"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
	"iot-monopoly/finance"
	"iot-monopoly/movement"
)

func Init() {
	eventing.StartExternalEventHandler()
	movement.StartEventHandler()
	finance.StartEventHandler()
}

//start with CompileDaemon -command="./iot-monopoly"
func main() {

	Init()
	//TODO get player infos from board
	playerOneId := uuid.New().String()
	playerTwoId := uuid.New().String()
	board.StartGame([]boardDomain.Player{{playerOneId, 0, 1000}, {playerTwoId, 0, 1000}})

	//TODO remove later
	fmt.Println("Sending external Events!!")
	events := []eventing.SensorEvent{
		{playerOneId, 5},
		{playerTwoId, 3},
		{playerOneId, 9},
		{playerTwoId, 7},
		{playerOneId, 15},
		{playerTwoId, 12},
		{playerOneId, 3},
		{playerTwoId, 1},
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
