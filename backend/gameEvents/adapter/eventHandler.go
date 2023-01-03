package gameEventsAdapter

import (
	"context"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing"
	"iot-monopoly/gameEvents/adapter/repository"
	domain "iot-monopoly/gameEvents/domain"
	playerDomain "iot-monopoly/player/domain"
	"strconv"
)

func StartEventListeners() {

	startGameStartedEventHandler()
	startPlayerMovedEventHandler()
}

func startGameStartedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			repository.InitFields()
		},
		Matcher: string(eventing.GAME_STARTED),
	})
}

func startPlayerMovedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			playerMovedEvent := e.Data.(*playerDomain.PlayerMovedEvent)

			eventField := repository.GetFieldById(strconv.FormatInt(int64(playerMovedEvent.FieldIndex), 10))

			switch eventField.Type() {

			case domain.DRAW_CARD:
				DrawCard(playerMovedEvent.PlayerId)
				break
			case domain.GOTO_PRISON:
				//TODO implement
				break
			case domain.PAY_TAX:
				eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, domain.NewGameEventWithFee("Bank", playerMovedEvent.PlayerId, 200))
				break

			}
		},
		Matcher: string(eventing.PLAYER_MOVED),
	})
}
