package game_master

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/eventing"
)

type Player struct {
	Id int
	// TODO maybe change position to field
	Position int
	Balance  int
}

type Field interface {
	OnPlayerEnter(player *Player)
}

//TODO consolidate Name to super class?
type PropertyField struct {
	Name  string
	Price int
	Owner *Player
}

type EventField struct {
	Name  string
	Event func(player *Player)
}

type BasicField struct {
	Name string
}

func (eventField EventField) OnPlayerEnter(player *Player) {

	eventField.Event(player)
}

func (_ BasicField) OnPlayerEnter(player *Player) {

	// do nothing
}

func (propertyField PropertyField) OnPlayerEnter(player *Player) {

	if propertyField.Owner == player {
		fmt.Println("player owns the property..")
	} else {
		fmt.Println("player does not own the property..")

	}
}

var players []Player
var fields []Field

func Init() {

	//TODO maybe make this an initial call from the GUI?
	fields = []Field{
		BasicField{"Start"},
		PropertyField{"Property purple 1", 100, nil},
		PropertyField{"Property purple 2", 100, nil},
		EventField{"Ereignisfeld 1", func(player *Player) {
			//TODO implement ereignis
			fmt.Println("Ereignisfeld")
		}},
		BasicField{"Gefaengnis"},
		EventField{"Ereignisfeld 2", func(player *Player) {
			//TODO implement ereignis
			fmt.Println("Ereignisfeld")
		}},
		PropertyField{"Property orange 1", 100, nil},
		PropertyField{"Property orange 2", 100, nil},
		BasicField{"Frei parken"},
		PropertyField{"Property green 1", 100, nil},
		EventField{"Start", func(player *Player) {
			//TODO remove money from bankAccount
			fmt.Println("Remove money from Bank account")
		}},
		PropertyField{"Property green 2", 100, nil},
		EventField{"Gehe ins gefaengnis", func(player *Player) {
			fmt.Println("Player has to go to prison")
			// TODO this field index for prison should not be magic
			movePlayer(player.Id, 4)
		}},
		PropertyField{"Property blue 1", 100, nil},
		EventField{"Ereignisfeld 3", func(player *Player) {
			//TODO implement ereignis
			fmt.Println("Ereignisfeld")
		}},
		PropertyField{"Property blue 2", 100, nil},
	}
	players = []Player{{1, 0, 1_000}, {2, 0, 1_000}}

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
			movePlayer(sensorEvent.PlayerId, sensorEvent.FieldIndex)
			printGameState()
		},
		func(err error) {
			fmt.Println(err)
		})
}

func printGameState() {

	fmt.Println("-------------------------------------")
	fmt.Println("Game State")
	fmt.Println("Players:")
	fmt.Println(players)
	fmt.Println("Fields:")
	fmt.Println(fields)
	fmt.Println("-------------------------------------")
}

func movePlayer(playerId int, fieldIndex int) {

	for i := range players {
		if players[i].Id == playerId {

			player := &players[i]

			//TODO improve this logic
			if (player.Position >= len(fields)-6 && player.Position < len(fields)) && (fieldIndex >= 0 && fieldIndex <= 5) {
				fmt.Println("Player completed a course, add money to balance")
				player.Balance += 100
			}

			fmt.Printf("Move player %d to fieldIndex %d\n", playerId, fieldIndex)
			player.Position = fieldIndex
			fields[fieldIndex].OnPlayerEnter(player)
		}
	}
}
