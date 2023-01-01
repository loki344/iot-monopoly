package player

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

func Init(playerCount int) ([]domain.Player, error) {
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

func MovePlayer(playerId string, fieldId int) error {

	player := getPlayer(playerId)
	player.SetPosition(fieldId)

	if currentPlayerIndex == len(players)-1 {
		currentPlayerIndex = 0
	} else {
		currentPlayerIndex++
	}
	currentPlayer = &players[currentPlayerIndex]
	return nil
}

func getPlayer(playerId string) *domain.Player {

	for i := range players {
		if players[i].Id() == playerId {
			return &players[i]
		}
	}

	panic(fmt.Sprintf("Player with id %s not found", playerId))
}

func GetPlayer(playerId string) *PlayerDTO {

	return dtoFromPlayer(getPlayer(playerId))
}

func GetCurrentPlayer() *PlayerDTO {
	return dtoFromPlayer(currentPlayer)
}

func GetPlayers() []*PlayerDTO {
	dtos := make([]*PlayerDTO, len(players))
	for i := range players {
		dtos[i] = dtoFromPlayer(&players[i])
	}
	return dtos
}

type PlayerDTO struct {
	Id        string `json:"id"`
	Position  int    `json:"position"`
	AccountId string `json:"accountId"`
}

func dtoFromPlayer(player *domain.Player) *PlayerDTO {
	return &PlayerDTO{player.Id(), player.Position(), player.AccountId()}
}
