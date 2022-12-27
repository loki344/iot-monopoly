package playerDomain

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

type PlayerMovedEvent struct {
	eventingDomain.BaseEvent
	PlayerId   string
	FieldIndex int
}

func NewPlayerMovedEvent(playerId string, fieldIndex int) *PlayerMovedEvent {
	return &PlayerMovedEvent{eventingDomain.EventType(&PlayerMovedEvent{}), playerId, fieldIndex}
}
