package gameEventsDomain

import eventingDomain "iot-monopoly/eventing/domain"

type GameEventWithPayout struct {
	eventingDomain.BaseEvent
	PlayerId string
	Amount   int
}

func NewGameEventWithPayout(playerId string, amount int) *GameEventWithPayout {

	return &GameEventWithPayout{BaseEvent: eventingDomain.EventType(&GameEventWithPayout{}), PlayerId: playerId, Amount: amount}
}

type GameEventWithFee struct {
	eventingDomain.BaseEvent
	PlayerId    string
	RecipientId string
	Fee         int
}

func NewGameEventWithFee(recipientId string, playerId string, fee int) *GameEventWithFee {

	return &GameEventWithFee{BaseEvent: eventingDomain.EventType(&GameEventWithFee{}), RecipientId: recipientId, PlayerId: playerId, Fee: fee}
}

type CardDrewEvent struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCardDrewEvent(title string, text string) *CardDrewEvent {
	return &CardDrewEvent{BaseEvent: eventingDomain.EventType(&CardDrewEvent{}), Title: title, Text: text}
}
