package domain

import eventingDomain "iot-monopoly/eventing/domain"

type GameEndedEvent struct {
	eventingDomain.BaseEvent
	WinnerId string
}

func NewGameEndedEvent(winnerId string) *GameEndedEvent {
	return &GameEndedEvent{BaseEvent: eventingDomain.EventType(&GameEndedEvent{}), WinnerId: winnerId}
}
