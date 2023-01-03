package repository

import (
	"errors"
	"fmt"
	domain "iot-monopoly/player/domain"
)

var players []domain.Player

//TODO use static uuid
var defaultPlayers = []domain.Player{
	*domain.NewPlayer("Player_1", 1, "Account_Player_1"),
	*domain.NewPlayer("Player_2", 1, "Account_Player_2"),
	*domain.NewPlayer("Player_3", 1, "Account_Player_3"),
	*domain.NewPlayer("Player_4", 1, "Account_Player_4"),
	//33-A8-8A-10 Card 1
	//	1304-B6-1A Card 2
	// 43-F1-E70E Card 3
	// A3-D9-350F Card 4
}

var currentPlayer *domain.Player
var currentPlayerIndex = 0

func InitPlayers(playerCount int) ([]domain.Player, error) {
	players = nil

	if playerCount < 1 || playerCount > 4 {
		errorMsg := fmt.Sprintf("invalid playerCount %d (must be between 1 and 4)", playerCount)
		fmt.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]domain.Player, playerCount)

	copy(newPlayers, defaultPlayers)
	players = newPlayers
	currentPlayer = &newPlayers[0]

	return players, nil
}

func GetPlayers() []domain.Player {
	return players
}

func FindPlayerById(playerId string) *domain.Player {

	for i := range players {
		if players[i].Id() == playerId {
			return &players[i]
		}
	}

	panic(fmt.Sprintf("Player with id %s not found", playerId))
}

func GetCurrentPlayer() *domain.Player {
	return currentPlayer
}

func GetCurrentPlayerIndex() int {
	return currentPlayerIndex
}

func SaveCurrentPlayerIndex(index int) {
	currentPlayerIndex = index
	currentPlayer = &players[index]
}
