package board

import (
	"errors"
	"fmt"
	"iot-monopoly/board/boardDomain"
	"iot-monopoly/eventing"
)

var Players []boardDomain.Player
var Fields []boardDomain.Field

var DefaultFields = []boardDomain.Field{
	boardDomain.BasicField{"Start"},
	boardDomain.PropertyField{"Property purple 1", 100, nil},
	boardDomain.PropertyField{"Property purple 2", 100, nil},
	boardDomain.EventField{"Ereignisfeld 1", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.BasicField{"Gefaengnis"},
	boardDomain.EventField{"Ereignisfeld 2", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.PropertyField{"Property orange 1", 100, nil},
	boardDomain.PropertyField{"Property orange 2", 100, nil},
	boardDomain.BasicField{"Frei parken"},
	boardDomain.PropertyField{"Property green 1", 100, nil},
	boardDomain.EventField{"Start", func(player *boardDomain.Player) {
		//TODO remove money from bankAccount
		fmt.Println("Remove money from Bank account")
	}},
	boardDomain.PropertyField{"Property green 2", 100, nil},
	boardDomain.EventField{"Gehe ins gefaengnis", func(player *boardDomain.Player) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		MovePlayer(player.Id, 4)
	}},
	boardDomain.PropertyField{"Property blue 1", 100, nil},
	boardDomain.EventField{"Ereignisfeld 3", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.PropertyField{"Property blue 2", 100, nil},
}

func StartGame(players []boardDomain.Player) {

	InitBoard(nil, players)
}

func InitBoard(fields []boardDomain.Field, players []boardDomain.Player) {

	if fields != nil {
		Fields = fields
	} else {
		fmt.Println("Initializing default fields")
		Fields = DefaultFields
	}
	Players = players
}

func MovePlayer(playerId string, fieldIndex int) error {

	if fieldIndex > len(Fields)-1 || fieldIndex < 0 {
		return errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", fieldIndex, len(Fields)))
	}

	player := GetPlayer(playerId)

	//TODO get rid of magic numbers 10!!
	if (player.Position >= 10 && player.Position < 16) && (fieldIndex >= 0 && fieldIndex <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, boardDomain.LapFinishedEvent{player.Id})
	}

	fmt.Printf("MovePlayer player %s to fieldIndex %d\n", player.Id, fieldIndex)
	player.Position = fieldIndex
	(*GetField(fieldIndex)).OnPlayerEnter(player)
	return nil
}

func GetPlayer(playerId string) *boardDomain.Player {

	for i := range Players {
		if Players[i].Id == playerId {
			return &Players[i]
		}
	}

	panic(fmt.Sprintf("Player with id %s not found", playerId))
}

func GetField(fieldIndex int) *boardDomain.Field {

	return &Fields[fieldIndex]
}
