package board

import (
	"errors"
	"fmt"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"strconv"
)

var players []boardDomain.Player

//TODO use static uuid
var defaultPlayers = []boardDomain.Player{
	{"Player_1", 0, "Account_Player_1"},
	{"Player_2", 0, "Account_Player_2"},
	{"Player_3", 0, "Account_Player_3"},
	{"Player_4", 0, "Account_Player_4"},
	//33-A8-8A-10 Card 1
	//	1304-B6-1A Card 2
	// 43-F1-E70E Card 3
	// A3-D9-350F Card 4
}

var currentPlayer *boardDomain.Player
var currentPlayerIndex = 0

func StartGame(playerCount int) ([]boardDomain.Player, error) {
	eventing.FireEvent(eventing.GAME_STARTED, boardDomain.NewGameStartedEvent(playerCount))
	players = nil
	initFields()

	if playerCount < 1 || playerCount > 4 {
		errorMsg := fmt.Sprintf("invalid playerCount %d (must be between 1 and 4)", playerCount)
		fmt.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]boardDomain.Player, playerCount)

	copy(newPlayers, defaultPlayers)
	players = newPlayers
	currentPlayer = &newPlayers[0]

	return players, nil
}

func MovePlayer(playerId string, fieldId int) error {

	totalFieldCount := GetFieldsCount()
	if fieldId > totalFieldCount-1 || fieldId < 0 {
		return errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", fieldId, totalFieldCount))
	}

	player := GetPlayer(playerId)
	if player.Position == fieldId {
		fmt.Println(fmt.Errorf("player already at position %d", fieldId))
		return nil
	}

	//TODO get rid of magic numbers 10!!
	if (player.Position >= 10 && player.Position < totalFieldCount) && (fieldId >= 0 && fieldId <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, boardDomain.NewLapFinishedEvent(player.Id))
	}

	fmt.Printf("MovePlayer player %s to fieldId %d\n", player.Id, fieldId)
	player.Position = fieldId
	GetFieldById(strconv.FormatInt(int64(fieldId), 10)).OnPlayerEnter(player)

	if currentPlayerIndex == len(players)-1 {
		currentPlayerIndex = 0
	} else {
		currentPlayerIndex++
	}
	currentPlayer = &players[currentPlayerIndex]
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

func GetCurrentPlayer() *boardDomain.Player {
	return currentPlayer
}

func GetPlayers() []boardDomain.Player {
	return players
}
