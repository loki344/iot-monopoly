package boardDomain

import (
	eventingDomain "iot-monopoly/eventing/domain"
)

type PlayerOnFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func NewPlayerOnUnownedFieldEvent(playerId string, property PropertyField) *PlayerOnFieldEvent {
	return &PlayerOnFieldEvent{eventingDomain.EventType(&PlayerOnFieldEvent{}), playerId, property}
}

type PlayerOnOwnedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	OwnerId  string
	Fee      int
}

func NewPlayerOnOwnedFieldEvent(playerId string, property PropertyField) *PlayerOnOwnedFieldEvent {
	return &PlayerOnOwnedFieldEvent{eventingDomain.EventType(&PlayerOnOwnedFieldEvent{}), playerId, property.OwnerId, property.GetPropertyFee()}
}

type PropertyTransferCreatedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
	ReceiverId    string
	SenderId      string
	Price         int
}

func NewPropertyTransferCreatedEvent(id string, receiverId string, senderId string, price int) *PropertyTransferCreatedEvent {
	return &PropertyTransferCreatedEvent{eventingDomain.EventType(&PropertyTransferCreatedEvent{}), id, receiverId, senderId, price}
}

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

type CreditAddedEvent struct {
	eventingDomain.BaseEvent
	RecipientAccountId string
	Amount             int
}

func NewCreditAddedEvent(recipientAccountId string, amount int) *CreditAddedEvent {

	return &CreditAddedEvent{BaseEvent: eventingDomain.EventType(&CreditAddedEvent{}), RecipientAccountId: recipientAccountId, Amount: amount}
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
	Title  string
	Text   string
	Action func(player *Player)
	Player *Player
}

func (card Card) TriggerAction() {
	card.Action(card.Player)
}

type CardDTO struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCard(title string, text string, action func(player *Player)) *Card {

	return &Card{Title: title, Text: text, Action: action}
}
