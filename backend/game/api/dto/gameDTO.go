package dto

import "iot-monopoly/game/domain"

type GameDTO struct {
	Players            []PlayerDTO `json:"players"`
	CurrentPlayerIndex int         `json:"currentPlayerIndex"`
	PlayerCount        int         `json:"playerCount"`
	Ended              bool        `json:"ended"`
}

func GameDTOFromGame(game *domain.Game) *GameDTO {
	return &GameDTO{
		Players:            PlayersDTOFromPlayers(game.Players()),
		CurrentPlayerIndex: game.CurrentPlayerIndex(),
		PlayerCount:        game.PlayerCount(),
		Ended:              game.Ended(),
	}
}
