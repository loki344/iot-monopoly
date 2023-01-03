package gameEventsAdapter

import (
	"iot-monopoly/gameEvents/adapter/repository"
)

func DrawCard(playerId string) {

	repository.GetNextCard(playerId)
}

func ConfirmCard() {

	card := repository.GetCurrentCard()
	card.TriggerAction()
	repository.DeleteCard(card)
}
