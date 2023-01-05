package events

import eventingDomain "iot-monopoly/eventing/domain"

type GameEventWithFeeAcceptedEvent struct {
	eventingDomain.BaseEvent
	PayerId     string
	RecipientId string
	Fee         int
}

func NewGameEventWithFeeAcceptedEvent(recipientId string, payerId string, fee int) *GameEventWithFeeAcceptedEvent {

	return &GameEventWithFeeAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithFeeAcceptedEvent{}), RecipientId: recipientId, PayerId: payerId, Fee: fee}
}
