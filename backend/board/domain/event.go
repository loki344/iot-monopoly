package boardDomain

import (
	eventingDomain "iot-monopoly/communication/domain"
)

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func NewGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}

type PlayerMovedEvent struct {
	eventingDomain.BaseEvent
	PlayerId   string
	FieldIndex int
}

func NewPlayerMovedEvent(playerId string, fieldIndex int) *PlayerMovedEvent {
	return &PlayerMovedEvent{eventingDomain.EventType(&PlayerMovedEvent{}), playerId, fieldIndex}
}
