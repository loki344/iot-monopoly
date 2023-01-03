package playerAdapter

import (
	"iot-monopoly/player/adapter/repository"
	domain "iot-monopoly/player/domain"
)

func MovePlayer(playerId string, fieldId int) error {

	player := repository.FindPlayerById(playerId)
	player.SetPosition(fieldId)

	currentPlayerIndex := repository.GetCurrentPlayerIndex()

	if currentPlayerIndex == len(GetPlayers())-1 {
		repository.SaveCurrentPlayerIndex(0)
	} else {
		repository.SaveCurrentPlayerIndex(currentPlayerIndex + 1)
	}

	return nil
}

func GetPlayer(playerId string) *PlayerDTO {

	return dtoFromPlayer(repository.FindPlayerById(playerId))
}

func GetCurrentPlayer() *PlayerDTO {
	return dtoFromPlayer(repository.GetCurrentPlayer())
}

func GetPlayers() []*PlayerDTO {
	players := repository.GetPlayers()
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
