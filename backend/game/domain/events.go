package domain

import eventingDomain "iot-monopoly/eventing/domain"

type CardDrewEvent struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCardDrewEvent(title string, text string) *CardDrewEvent {
	return &CardDrewEvent{BaseEvent: eventingDomain.EventType(&CardDrewEvent{}), Title: title, Text: text}
}

type GameEndedEvent struct {
	eventingDomain.BaseEvent
	WinnerId string
}

func NewGameEndedEvent(winnerId string) *GameEndedEvent {
	return &GameEndedEvent{BaseEvent: eventingDomain.EventType(&GameEndedEvent{}), WinnerId: winnerId}
}

type GameEventWithFeeAcceptedEvent struct {
	eventingDomain.BaseEvent
	PayerId     string
	RecipientId string
	Fee         int
}

func NewGameEventWithFeeAcceptedEvent(recipientId string, payerId string, fee int) *GameEventWithFeeAcceptedEvent {

	return &GameEventWithFeeAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithFeeAcceptedEvent{}), RecipientId: recipientId, PayerId: payerId, Fee: fee}
}

type GameEventWithPayoutAcceptedEvent struct {
	eventingDomain.BaseEvent
	RecipientId string
	Amount      int
}

func NewGameEventWithPayoutAcceptedEvent(recipientId string, amount int) *GameEventWithPayoutAcceptedEvent {

	return &GameEventWithPayoutAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithPayoutAcceptedEvent{}), RecipientId: recipientId, Amount: amount}
}

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func NewGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}

type PlayerOnOwnedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	OwnerId  string
	Fee      int
}

func NewPlayerOnOwnedFieldEvent(playerId string, ownerId string, fee int) *PlayerOnOwnedFieldEvent {
	return &PlayerOnOwnedFieldEvent{eventingDomain.EventType(&PlayerOnOwnedFieldEvent{}), playerId, ownerId, fee}
}

type PlayerOnUnownedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func NewPlayerOnUnownedFieldEvent(playerId string, property *PropertyField) *PlayerOnUnownedFieldEvent {
	return &PlayerOnUnownedFieldEvent{eventingDomain.EventType(&PlayerOnUnownedFieldEvent{}), playerId, *property}
}

type TransactionCreatedEvent struct {
	eventingDomain.BaseEvent
	RecipientId string
	SenderId    string
	Amount      int
}

func NewTransactionCreatedEvent(recipientId string, senderId string, amount int) TransactionCreatedEvent {

	return TransactionCreatedEvent{BaseEvent: eventingDomain.EventType(&TransactionCreatedEvent{}), RecipientId: recipientId, SenderId: senderId, Amount: amount}
}

type TransactionResolvedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
}

func NewTransactionResolvedEvent(transactionId string) TransactionResolvedEvent {

	return TransactionResolvedEvent{BaseEvent: eventingDomain.EventType(&TransactionResolvedEvent{}), TransactionId: transactionId}
}
