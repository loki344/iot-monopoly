package gameAdapter

import (
	"iot-monopoly/communication"
	domain "iot-monopoly/game/domain"
)

func StartGame(playerCount int) {

	//Here is the entry point to implement the logic to connect multiple players online etc.
	communication.FireEvent(communication.GAME_STARTED, domain.NewGameStartedEvent(playerCount))
}

func EndGame(status string) {
	communication.FireEvent(communication.GAME_ENDED, domain.NewGameEndedEvent(status))
}
