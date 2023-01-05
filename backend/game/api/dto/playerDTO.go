package dto

import "iot-monopoly/game/domain"

type PlayerDTO struct {
	Id                        string     `json:"id"`
	Position                  int        `json:"position"`
	Account                   AccountDTO `json:"account"`
	EscapeFromPrisonCardCount int        `json:"escapeFromPrisonCardCount"`
}

func PlayersDTOFromPlayers(players []domain.Player) []PlayerDTO {

	playerDTOs := make([]PlayerDTO, len(players))
	for i := range playerDTOs {
		playerDTOs[i] = PlayerDTOFromPlayer(players[i])
	}
	return playerDTOs
}

func PlayerDTOFromPlayer(player domain.Player) PlayerDTO {
	return PlayerDTO{Id: player.Id(), Position: player.Position(), Account: AccountDTOFromAccount(player.Account()), EscapeFromPrisonCardCount: player.EscapeFromPrisonCardCount()}
}
