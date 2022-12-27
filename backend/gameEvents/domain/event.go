package gameEventsDomain

import eventingDomain "iot-monopoly/communication/domain"

type CardWithPayoutEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Amount   int
}

func NewCardWithPayoutDrewEvent(playerId string, amount int) *CardWithPayoutEvent {

	return &CardWithPayoutEvent{BaseEvent: eventingDomain.EventType(&CardWithPayoutEvent{}), PlayerId: playerId, Amount: amount}
}

type CardWithFeeEvent struct {
	eventingDomain.BaseEvent
	PlayerId    string
	RecipientId string
	Fee         int
}

func NewCardWithFeeDrewEvent(recipientId string, playerId string, fee int) *CardWithFeeEvent {

	return &CardWithFeeEvent{BaseEvent: eventingDomain.EventType(&CardWithFeeEvent{}), RecipientId: recipientId, PlayerId: playerId, Fee: fee}
}

type CardDrewEvent struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCardDrewEvent(title string, text string) *CardDrewEvent {
	return &CardDrewEvent{BaseEvent: eventingDomain.EventType(&CardDrewEvent{}), Title: title, Text: text}
}
