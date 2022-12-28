package player

import (
	"errors"
	"fmt"
	boardDomain "iot-monopoly/player/domain"
)

var players []*boardDomain.Player

//TODO use static uuid
var defaultPlayers = []*boardDomain.Player{
	boardDomain.NewPlayer("Player_1", 0, "Account_Player_1"),
	boardDomain.NewPlayer("Player_2", 0, "Account_Player_2"),
	boardDomain.NewPlayer("Player_3", 0, "Account_Player_3"),
	boardDomain.NewPlayer("Player_4", 0, "Account_Player_4"),
	//33-A8-8A-10 Card 1
	//	1304-B6-1A Card 2
	// 43-F1-E70E Card 3
	// A3-D9-350F Card 4
}

var currentPlayer *boardDomain.Player
var currentPlayerIndex = 0

func Init(playerCount int) ([]*boardDomain.Player, error) {
	players = nil

	if playerCount < 1 || playerCount > 4 {
		errorMsg := fmt.Sprintf("invalid playerCount %d (must be between 1 and 4)", playerCount)
		fmt.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]*boardDomain.Player, playerCount)

	copy(newPlayers, defaultPlayers)
	players = newPlayers
	currentPlayer = newPlayers[0]

	return players, nil
}

func MovePlayer(playerId string, fieldId int) error {

	player := GetPlayer(playerId)
	player.SetPosition(fieldId)

	if currentPlayerIndex == len(players)-1 {
		currentPlayerIndex = 0
	} else {
		currentPlayerIndex++
	}
	currentPlayer = players[currentPlayerIndex]
	return nil
}

func GetPlayer(playerId string) *boardDomain.Player {

	for i := range players {
		if players[i].Id() == playerId {
			return players[i]
		}
	}

	panic(fmt.Sprintf("Player with id %s not found", playerId))
}

func GetCurrentPlayer() *boardDomain.Player {
	return currentPlayer
}

func GetPlayers() []*boardDomain.Player {
	return players
}
