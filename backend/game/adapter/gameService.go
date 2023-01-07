package adapter

import (
	"iot-monopoly/game/adapter/repository"
	"iot-monopoly/game/api/dto"
	"iot-monopoly/game/domain"
)

func NewGame(playerCount int) *domain.Game {
	return repository.SaveGame(domain.NewGame(playerCount))
}

func GetCurrentGame() *domain.Game {
	return repository.GetCurrentGame()
}

func BuyProperty(propertyIndex int, buyerId string) {
	referenceId := repository.GetCurrentGame().BuyProperty(propertyIndex, buyerId)
	transaction := domain.NewTransaction(referenceId, "Bank", GetCurrentGame().GetPlayerById(buyerId).Account().Id(), repository.FindPropertyByIndex(propertyIndex).GetPrice())
	repository.CreateTransaction(transaction)
}

func ConfirmCurrentCard() {
	GetCurrentGame().ConfirmCurrentCard()
}

func EndGame() {

	GetCurrentGame().End("")
	repository.DeleteGame()
}

func FindPlayerById(playerId string) dto.PlayerDTO {
	propertyCount := 0
	for _, property := range GetCurrentGame().Properties() {
		if property.OwnerId() == playerId {
			propertyCount++
		}
	}

	return dto.PlayerDTOFromPlayer(*GetCurrentGame().GetPlayerById(playerId), propertyCount)
}
