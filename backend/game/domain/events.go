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
	PlayerId           string
	PropertyName       string
	PropertyIndex      int
	PropertyPrice      int
	HousePrice         int
	HotelPrice         int
	RevenueNormal      int
	RevenueOneHouse    int
	RevenueTwoHouses   int
	RevenueThreeHouses int
	RevenueFourHouses  int
	RevenueHotel       int
}

func newPlayerOnUnownedFieldEvent(playerId string, property *PropertyField) *PlayerOnUnownedFieldEvent {
	fd := property.FinancialDetails()
	return &PlayerOnUnownedFieldEvent{BaseEvent: eventingDomain.EventType(&PlayerOnUnownedFieldEvent{}), PlayerId: playerId, PropertyIndex: property.index,
		PropertyName: property.Name(), PropertyPrice: fd.PropertyPrice(), HousePrice: fd.HousePrice(), HotelPrice: fd.HotelPrice(), RevenueNormal: fd.Revenue().Normal(),
		RevenueOneHouse: fd.Revenue().OneHouse(), RevenueTwoHouses: fd.Revenue().TwoHouses(), RevenueThreeHouses: fd.Revenue().ThreeHouses(),
		RevenueFourHouses: fd.Revenue().FourHouses(), RevenueHotel: fd.Revenue().Hotel()}
}

type TransactionCreatedEvent struct {
	eventingDomain.BaseEvent
	Id          string
	RecipientId string
	SenderId    string
	Amount      int
}

func newTransactionCreatedEvent(id string, recipientId string, senderId string, amount int) TransactionCreatedEvent {

	return TransactionCreatedEvent{BaseEvent: eventingDomain.EventType(&TransactionCreatedEvent{}), Id: id, RecipientId: recipientId, SenderId: senderId, Amount: amount}
}

type TransactionResolvedEvent struct {
	eventingDomain.BaseEvent
	TransactionId string
}

func newTransactionResolvedEvent(transactionId string) TransactionResolvedEvent {

	return TransactionResolvedEvent{BaseEvent: eventingDomain.EventType(&TransactionResolvedEvent{}), TransactionId: transactionId}
}

type PlayerDataUpdatedEvent struct {
	eventingDomain.BaseEvent
}

func NewPlayerDataUpdatedEvent() *PlayerDataUpdatedEvent {
	return &PlayerDataUpdatedEvent{BaseEvent: eventingDomain.EventType(&PlayerDataUpdatedEvent{})}
}

type AccountDataUpdatedEvent struct {
	eventingDomain.BaseEvent
}

func NewAccountDataUpdatedEvent() *AccountDataUpdatedEvent {
	return &AccountDataUpdatedEvent{BaseEvent: eventingDomain.EventType(&AccountDataUpdatedEvent{})}
}
