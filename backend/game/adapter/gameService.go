package adapter

import (
	"iot-monopoly/game/adapter/repository"
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
