package repository

import "iot-monopoly/game/domain"

func FindAccountById(accountId string) *domain.Account {

	if accountId == "Bank" {
		return GetCurrentGame().Bank().Account()
	}

	players := GetCurrentGame().Players()
	for i := range players {
		if players[i].Account().Id() == accountId {
			return players[i].Account()
		}
	}
	return nil
}
