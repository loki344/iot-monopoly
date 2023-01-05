package domain

import (
	eventingDomain "iot-monopoly/eventing/domain"
)

type PlayerOnOwnedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	OwnerId  string
	Fee      int
}

func NewPlayerOnOwnedFieldEvent(playerId string, ownerId string, fee int) *PlayerOnOwnedFieldEvent {
	return &PlayerOnOwnedFieldEvent{eventingDomain.EventType(&PlayerOnOwnedFieldEvent{}), playerId, ownerId, fee}
}
