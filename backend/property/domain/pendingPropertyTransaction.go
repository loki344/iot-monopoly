package domain

import "iot-monopoly/communication"

type PendingPropertyTransaction struct {
	id         string
	propertyId string
	buyerId    string
}

func (p PendingPropertyTransaction) Id() string {
	return p.id
}

func (p PendingPropertyTransaction) PropertyId() string {
	return p.propertyId
}

func (p PendingPropertyTransaction) BuyerId() string {
	return p.buyerId
}

func NewPendingPropertyTransaction(transactionId string, propertyId string, buyerId string, price int) *PendingPropertyTransaction {
	communication.FireEvent(communication.PROPERTY_TRANSFER_CREATED, NewPropertyTransferCreatedEvent(transactionId, "Bank", buyerId, price))
	return &PendingPropertyTransaction{id: transactionId, propertyId: propertyId, buyerId: buyerId}
}
