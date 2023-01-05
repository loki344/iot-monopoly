package events

import (
	eventingDomain "iot-monopoly/eventing/domain"
	"iot-monopoly/game/domain"
)

type PlayerOnUnownedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property domain.PropertyField
}

func NewPlayerOnUnownedFieldEvent(playerId string, property *domain.PropertyField) *PlayerOnUnownedFieldEvent {
	return &PlayerOnUnownedFieldEvent{eventingDomain.EventType(&PlayerOnUnownedFieldEvent{}), playerId, *property}
}
