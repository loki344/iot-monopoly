package adapter

import (
	"context"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	gameDomain "iot-monopoly/game/domain"
	"iot-monopoly/player/adapter/repository"
)

func StartEventListeners() {

	startGameStartedEventHandler()
}
func startGameStartedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			gameStartedEvent := e.Data.(*gameDomain.GameStartedEvent)

			repository.InitPlayers(gameStartedEvent.PlayerCount)
		},
		Matcher: string(eventing.GAME_STARTED),
	})
}
