package gameEvents

import (
	"context"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/communication"
	boardDomain "iot-monopoly/player/domain"
	"strconv"
)

func StartEventListeners() {

	startGameStartedEventHandler()
	startPlayerMovedEventHandler()
}

func startGameStartedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			initFields()
		},
		Matcher: string(communication.GAME_STARTED),
	})
}

func startPlayerMovedEventHandler() {

	communication.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			playerMovedEvent := e.Data.(boardDomain.PlayerMovedEvent)

			eventField := GetFieldById(strconv.FormatInt(int64(playerMovedEvent.FieldIndex), 10))

			if eventField != nil {
				eventField.OnPlayerEnter(playerMovedEvent.PlayerId)
			}
		},
		Matcher: string(communication.PLAYER_MOVED),
	})
}
