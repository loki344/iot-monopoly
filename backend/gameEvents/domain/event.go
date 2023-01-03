package gameEventsDomain

import eventingDomain "iot-monopoly/eventing/domain"

type GameEventWithPayoutAcceptedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Amount   int
}

func NewGameEventWithPayoutAcceptedEvent(playerId string, amount int) *GameEventWithPayoutAcceptedEvent {

	return &GameEventWithPayoutAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithPayoutAcceptedEvent{}), PlayerId: playerId, Amount: amount}
}

type GameEventWithFeeAcceptedEvent struct {
	eventingDomain.BaseEvent
	PlayerId    string
	RecipientId string
	Fee         int
}

func NewGameEventWithFeeAcceptedEvent(recipientId string, playerId string, fee int) *GameEventWithFeeAcceptedEvent {

	return &GameEventWithFeeAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithFeeAcceptedEvent{}), RecipientId: recipientId, PlayerId: playerId, Fee: fee}
}

type CardDrewEvent struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCardDrewEvent(title string, text string) *CardDrewEvent {
	return &CardDrewEvent{BaseEvent: eventingDomain.EventType(&CardDrewEvent{}), Title: title, Text: text}
}
