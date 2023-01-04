package repository

import "iot-monopoly/game/domain"

var currentGame *domain.Game

func SaveGame(game *domain.Game) *domain.Game {
	currentGame = game
	return currentGame
}

func GetCurrentGame() *domain.Game {
	return currentGame
}
