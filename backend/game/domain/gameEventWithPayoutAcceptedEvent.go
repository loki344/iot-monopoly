package domain

import eventingDomain "iot-monopoly/eventing/domain"

type GameEventWithPayoutAcceptedEvent struct {
	eventingDomain.BaseEvent
	RecipientId string
	Amount      int
}

func NewGameEventWithPayoutAcceptedEvent(recipientId string, amount int) *GameEventWithPayoutAcceptedEvent {

	return &GameEventWithPayoutAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithPayoutAcceptedEvent{}), RecipientId: recipientId, Amount: amount}
}
