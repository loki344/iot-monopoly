package domain

import (
	eventingDomain "iot-monopoly/eventing/domain"
)

type PlayerOnUnownedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func NewPlayerOnUnownedFieldEvent(playerId string, property *PropertyField) *PlayerOnUnownedFieldEvent {
	return &PlayerOnUnownedFieldEvent{eventingDomain.EventType(&PlayerOnUnownedFieldEvent{}), playerId, *property}
}
