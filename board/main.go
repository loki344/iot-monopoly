package board

import (
	"fmt"
	"iot-monopoly/board/boardDomain"
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
		//TODO trigger movement of player
	}},
	boardDomain.PropertyField{"Property blue 1", 100, nil},
	boardDomain.EventField{"Ereignisfeld 3", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.PropertyField{"Property blue 2", 100, nil},
}

func InitBoard(fields []boardDomain.Field, players []boardDomain.Player) {

	if fields != nil {
		Fields = fields
	} else {
		Fields = DefaultFields
	}
	Players = players
}

func GetPlayer(playerId int) *boardDomain.Player {

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
