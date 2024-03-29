package dto

import "iot-monopoly/game/domain"

type PlayerDTO struct {
	Id                        string     `json:"id"`
	Position                  int        `json:"position"`
	Account                   AccountDTO `json:"account"`
	EscapeFromPrisonCardCount int        `json:"escapeFromPrisonCardCount"`
	PropertyCount             int        `json:"propertyCount"`
}

func PlayersDTOFromPlayers(players []domain.Player, propertyCounts map[string]int) []PlayerDTO {

	playerDTOs := make([]PlayerDTO, len(players))
	for i := range playerDTOs {
		playerDTOs[i] = PlayerDTOFromPlayer(players[i], propertyCounts[players[i].Id()])
	}
	return playerDTOs
}

func PlayerDTOFromPlayer(player domain.Player, propertyCount int) PlayerDTO {
	return PlayerDTO{Id: player.Id(), Position: player.Position(), Account: AccountDTOFromAccount(player.Account()), EscapeFromPrisonCardCount: player.EscapeFromPrisonCardCount(), PropertyCount: propertyCount}
}
