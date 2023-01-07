package dto

import "iot-monopoly/game/domain"

type GameDTO struct {
	Players            []PlayerDTO `json:"players"`
	CurrentPlayerIndex int         `json:"currentPlayerIndex"`
	PlayerCount        int         `json:"playerCount"`
	Ended              bool        `json:"ended"`
}

func GameDTOFromGame(game *domain.Game) *GameDTO {

	propertyCountPerPlayerId := make(map[string]int)

	for _, player := range game.Players() {
		propertyCountPerPlayerId[player.Id()] = 0
	}

	for _, property := range game.Properties() {
		if property.OwnerId() != "" {
			propertyCountPerPlayerId[property.OwnerId()] = propertyCountPerPlayerId[property.OwnerId()] + 1
		}
	}

	return &GameDTO{
		Players:            PlayersDTOFromPlayers(game.Players(), propertyCountPerPlayerId),
		CurrentPlayerIndex: game.CurrentPlayerIndex(),
		PlayerCount:        game.PlayerCount(),
		Ended:              game.Ended(),
	}
}
