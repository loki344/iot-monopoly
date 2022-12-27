package game

import (
	"iot-monopoly/communication"
	gameDomain "iot-monopoly/game/domain"
)

func StartGame(playerCount int) {

	//Here is the entry point to implement the logic to connect multiple players online etc.
	communication.FireEvent(communication.GAME_STARTED, gameDomain.NewGameStartedEvent(playerCount))
}
