package events

import eventingDomain "iot-monopoly/eventing/domain"

type PlayerBankruptEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewPlayerBankruptEvent(playerId string) *PlayerBankruptEvent {
	return &PlayerBankruptEvent{BaseEvent: eventingDomain.EventType(&PlayerBankruptEvent{}), PlayerId: playerId}
}
