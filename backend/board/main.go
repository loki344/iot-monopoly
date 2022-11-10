package board

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
)

var players []boardDomain.Player
var fields []boardDomain.Field

//TODO use static uuid
var defaultPlayers = []boardDomain.Player{
	{"AA-19-F6-5B", 0, 1_000},
	{"53-C8-A4-28-110001", 0, 1_000},
	{"53-D8-A6-28-110001", 0, 1_000},
	{"53-81-A4-28-110001", 0, 1_000},
	//33-A8-8A-10 Card 1
	//	1304-B6-1A Card 2
	// 43-F1-E70E Card 3
	// A3-D9-350F Card 4
}

var tempFinancialDetails = boardDomain.FinancialDetails{100, 100, 100, boardDomain.Revenue{100, 200, 300, 400, 500, 800}}

var DefaultFields = []boardDomain.Field{
	boardDomain.BasicField{uuid.New().String(), "Start"},
	boardDomain.NewPropertyField("Property purple 1", uuid.NewString(), tempFinancialDetails),
	boardDomain.NewPropertyField("Property purple 2", uuid.NewString(), tempFinancialDetails),
	boardDomain.EventField{uuid.New().String(), "Ereignisfeld 1", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.BasicField{uuid.New().String(), "Gefaengnis"},
	boardDomain.EventField{uuid.New().String(), "Ereignisfeld 2", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.NewPropertyField("Property orange 1", uuid.NewString(), tempFinancialDetails),
	boardDomain.NewPropertyField("Property orange 2", uuid.NewString(), tempFinancialDetails),
	boardDomain.BasicField{uuid.New().String(), "Frei parken"},
	boardDomain.NewPropertyField("Property green 1", uuid.NewString(), tempFinancialDetails),
	boardDomain.EventField{uuid.New().String(), "Start", func(player *boardDomain.Player) {
		fmt.Printf("Remove 100 from Bank account of player %s\n", player.Id)
		player.Balance -= 100
	}},
	boardDomain.NewPropertyField("Property green 2", uuid.NewString(), tempFinancialDetails),
	boardDomain.EventField{uuid.New().String(), "Gehe ins gefaengnis", func(player *boardDomain.Player) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		MovePlayer(player.Id, 4)
	}},
	boardDomain.NewPropertyField("Property blue 1", uuid.NewString(), tempFinancialDetails),
	boardDomain.EventField{uuid.New().String(), "Ereignisfeld 3", func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
	boardDomain.NewPropertyField("Property blue 2", uuid.NewString(), tempFinancialDetails),
}

func StartGame(playerCount int) ([]boardDomain.Player, error) {
	players = nil
	fields = nil

	if playerCount < 1 || playerCount > 4 {
		errorMsg := fmt.Sprintf("invalid playerCount %d (must be between 1 and 4)", playerCount)
		fmt.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]boardDomain.Player, playerCount)

	copy(newPlayers, defaultPlayers)

	players := newPlayers[0:playerCount]

	initBoard(nil, players)
	return players, nil
}

func initBoard(initFields []boardDomain.Field, initPlayers []boardDomain.Player) {

	if initFields != nil {
		fields = initFields
	} else {
		fmt.Println("Initializing default initFields")
		fields = DefaultFields
	}
	players = initPlayers
}

func MovePlayer(playerId string, fieldIndex int) error {

	if fieldIndex > len(fields)-1 || fieldIndex < 0 {
		return errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", fieldIndex, len(fields)))
	}

	player := GetPlayer(playerId)
	if player.Position == fieldIndex {
		fmt.Println(fmt.Errorf("player already at position %d", fieldIndex))
		return nil
	}

	//TODO get rid of magic numbers 10!!
	if (player.Position >= 10 && player.Position < 16) && (fieldIndex >= 0 && fieldIndex <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, boardDomain.NewLapFinishedEvent(player.Id))
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
