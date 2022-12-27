package player

import (
	"context"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/communication"
	gameDomain "iot-monopoly/game/domain"
)

func StartEventListeners() {

	startGameStartedEventHandler()
}
func startGameStartedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			gameStartedEvent := e.Data.(*gameDomain.GameStartedEvent)

			Init(gameStartedEvent.PlayerCount)
		},
		Matcher: string(communication.GAME_STARTED),
	})
}
