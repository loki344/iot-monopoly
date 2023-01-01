package propertyDomain

import (
	eventingDomain "iot-monopoly/communication/domain"
)

type PlayerOnUnownedFieldEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func NewPlayerOnUnownedFieldEvent(playerId string, property PropertyField) *PlayerOnUnownedFieldEvent {
	return &PlayerOnUnownedFieldEvent{eventingDomain.EventType(&PlayerOnUnownedFieldEvent{}), playerId, property}
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
