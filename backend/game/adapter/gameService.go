package adapter

import (
	"iot-monopoly/eventing"
	domain "iot-monopoly/game/domain"
)

func StartGame(playerCount int) {

	//Here is the entry point to implement the logic to connect multiple players online etc.
	eventing.FireEvent(eventing.GAME_STARTED, domain.NewGameStartedEvent(playerCount))
}

func EndGame(status string) {
	eventing.FireEvent(eventing.GAME_ENDED, domain.NewGameEndedEvent(status))
}
