package events

import eventingDomain "iot-monopoly/eventing/domain"

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func NewGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}
