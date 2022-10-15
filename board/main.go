package board

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"iot-monopoly/eventing/domain"
)

var players []boardDomain.Player
var fields []boardDomain.Field

var DefaultFields = []boardDomain.Field{
	boardDomain.BasicField{uuid.New().String(), "Start"},
	boardDomain.PropertyField{uuid.New().String(), "Property purple 1", 100, ""},
	boardDomain.PropertyField{uuid.New().String(), "Property purple 2", 100, ""},
	boardDomain.EventField{uuid.New().String(), "Ereignisfeld 1", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.BasicField{uuid.New().String(), "Gefaengnis"},
	boardDomain.EventField{uuid.New().String(), "Ereignisfeld 2", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.PropertyField{uuid.New().String(), "Property orange 1", 100, ""},
	boardDomain.PropertyField{uuid.New().String(), "Property orange 2", 100, ""},
	boardDomain.BasicField{uuid.New().String(), "Frei parken"},
	boardDomain.PropertyField{uuid.New().String(), "Property green 1", 100, ""},
	boardDomain.EventField{uuid.New().String(), "Start", func(player *boardDomain.Player) {
		//TODO remove money from bankAccount
		fmt.Println("Remove money from Bank account")
	}},
	boardDomain.PropertyField{uuid.New().String(), "Property green 2", 100, ""},
	boardDomain.EventField{uuid.New().String(), "Gehe ins gefaengnis", func(player *boardDomain.Player) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		MovePlayer(player.Id, 4)
	}},
	boardDomain.PropertyField{uuid.New().String(), "Property blue 1", 100, ""},
	boardDomain.EventField{uuid.New().String(), "Ereignisfeld 3", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.PropertyField{uuid.New().String(), "Property blue 2", 100, ""},
}

func StartGame(players []boardDomain.Player) {

	players = nil
	fields = nil
	InitBoard(nil, players)
}

func InitBoard(initFields []boardDomain.Field, players []boardDomain.Player) {

	if initFields != nil {
		fields = initFields
	} else {
		fmt.Println("Initializing default initFields")
		fields = DefaultFields
	}
	players = players
}

func MovePlayer(playerId string, fieldIndex int) error {

	if fieldIndex > len(fields)-1 || fieldIndex < 0 {
		return errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", fieldIndex, len(fields)))
	}

	player := GetPlayer(playerId)

	//TODO get rid of magic numbers 10!!
	if (player.Position >= 10 && player.Position < 16) && (fieldIndex >= 0 && fieldIndex <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, eventingDomain.LapFinishedEvent{PlayerId: player.Id})
	}

	fmt.Printf("MovePlayer player %s to fieldIndex %d\n", player.Id, fieldIndex)
	player.Position = fieldIndex
	(*GetField(fieldIndex)).OnPlayerEnter(player)
	return nil
}

func GetPlayer(playerId string) *boardDomain.Player {

	for i := range players {
		if players[i].Id == playerId {
			return &players[i]
		}
	}

	panic(fmt.Sprintf("Player with id %s not found", playerId))
}

func GetField(fieldIndex int) *boardDomain.Field {

	return &fields[fieldIndex]
}

func GetFieldsCount() int {
	return len(fields)
}

func GetPlayers() []boardDomain.Player {
	return players
}

func GetFields() []boardDomain.Field {
	return fields
}
