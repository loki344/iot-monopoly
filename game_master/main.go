package game_master

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/eventing"
)

type Player struct {
	Id          int
	Position    *Field
	BankAccount *BankAccount
}

type Field struct {
}

type Property struct {
	Field
	Price int
	Owner *Player
}

type BankAccount struct {
	Owner   *Player
	Balance int
}

var players []Player
var fields []Field

func Init() {

	//TODO maybe make this an initial call from the GUI?

	//TODO move to other package?
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
			sensorEvent := msg.Payload.(eventing.SensorEvent)
			fmt.Println(sensorEvent)
			movePlayer(sensorEvent.PlayerId, sensorEvent.FieldId)
		},
		func(err error) {
			fmt.Println(err)
		})
}

func movePlayer(playerId int, fieldId int) {

}
