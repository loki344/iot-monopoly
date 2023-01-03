package domain

import "iot-monopoly/eventing"

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
	eventing.FireEvent(eventing.PROPERTY_TRANSFER_CREATED, NewPropertyTransferCreatedEvent(transactionId, "Bank", buyerId, price))
	return &PendingPropertyTransaction{id: transactionId, propertyId: propertyId, buyerId: buyerId}
}
