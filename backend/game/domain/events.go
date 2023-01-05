package domain

import eventingDomain "iot-monopoly/eventing/domain"

type CardDrewEvent struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func newCardDrewEvent(title string, text string) *CardDrewEvent {
	return &CardDrewEvent{BaseEvent: eventingDomain.EventType(&CardDrewEvent{}), Title: title, Text: text}
}

type GameEndedEvent struct {
	eventingDomain.BaseEvent
	WinnerId string
}

func newGameEndedEvent(winnerId string) *GameEndedEvent {
	return &GameEndedEvent{BaseEvent: eventingDomain.EventType(&GameEndedEvent{}), WinnerId: winnerId}
}

type GameEventWithFeeAcceptedEvent struct {
	eventingDomain.BaseEvent
	PayerId     string
	RecipientId string
	Fee         int
}

func newGameEventWithFeeAcceptedEvent(recipientId string, payerId string, fee int) *GameEventWithFeeAcceptedEvent {

	return &GameEventWithFeeAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithFeeAcceptedEvent{}), RecipientId: recipientId, PayerId: payerId, Fee: fee}
}

type GameEventWithPayoutAcceptedEvent struct {
	eventingDomain.BaseEvent
	RecipientId string
	Amount      int
}

func newGameEventWithPayoutAcceptedEvent(recipientId string, amount int) *GameEventWithPayoutAcceptedEvent {

	return &GameEventWithPayoutAcceptedEvent{BaseEvent: eventingDomain.EventType(&GameEventWithPayoutAcceptedEvent{}), RecipientId: recipientId, Amount: amount}
}

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func newGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func newLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}

type PlayerOnOwnedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	OwnerId  string
	Fee      int
}

func newPlayerOnOwnedFieldEvent(playerId string, ownerId string, fee int) *PlayerOnOwnedFieldEvent {
	return &PlayerOnOwnedFieldEvent{eventingDomain.EventType(&PlayerOnOwnedFieldEvent{}), playerId, ownerId, fee}
}

type PlayerOnUnownedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func newPlayerOnUnownedFieldEvent(playerId string, property *PropertyField) *PlayerOnUnownedFieldEvent {
	return &PlayerOnUnownedFieldEvent{eventingDomain.EventType(&PlayerOnUnownedFieldEvent{}), playerId, *property}
}

type TransactionCreatedEvent struct {
	eventingDomain.BaseEvent
	RecipientId string
	SenderId    string
	Amount      int
}

func newTransactionCreatedEvent(recipientId string, senderId string, amount int) TransactionCreatedEvent {

	return TransactionCreatedEvent{BaseEvent: eventingDomain.EventType(&TransactionCreatedEvent{}), RecipientId: recipientId, SenderId: senderId, Amount: amount}
}

type TransactionResolvedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
}

func newTransactionResolvedEvent(transactionId string) TransactionResolvedEvent {

	return TransactionResolvedEvent{BaseEvent: eventingDomain.EventType(&TransactionResolvedEvent{}), TransactionId: transactionId}
}
