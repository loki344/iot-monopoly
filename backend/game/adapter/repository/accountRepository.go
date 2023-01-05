package repository

import "iot-monopoly/game/domain"

func FindAccountById(accountId string) *domain.Account {

	players := GetCurrentGame().Players()
	for i := range players {
		if players[i].Account().Id() == accountId {
			return players[i].Account()
		}
	}
	return nil
}
