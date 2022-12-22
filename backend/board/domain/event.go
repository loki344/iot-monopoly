package boardDomain

import (
	eventingDomain "iot-monopoly/communication/domain"
)

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func NewGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}

type PlayerMovedEvent struct {
	eventingDomain.BaseEvent
	PlayerId   string
	FieldIndex int
}

func NewPlayerMovedEvent(playerId string, fieldIndex int) *PlayerMovedEvent {
	return &PlayerMovedEvent{eventingDomain.EventType(&PlayerMovedEvent{}), playerId, fieldIndex}
}

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

type Card struct {
	Title    string
	Text     string
	Action   func(playerId string)
	PlayerId string
}

func (card Card) TriggerAction() {
	card.Action(card.PlayerId)
}

type CardDTO struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCard(title string, text string, action func(playerId string)) *Card {

	return &Card{Title: title, Text: text, Action: action}
}
