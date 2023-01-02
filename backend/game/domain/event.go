package gameDomain

import eventingDomain "iot-monopoly/communication/domain"

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func NewGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}

type GameEndedEvent struct {
	eventingDomain.BaseEvent
	Status string
}

func NewGameEndedEvent(status string) *GameEndedEvent {
	return &GameEndedEvent{BaseEvent: eventingDomain.EventType(&GameEndedEvent{}), Status: status}
}
