package propertyDomain

import "iot-monopoly/communication"

type PendingPropertyTransaction struct {
	TransactionId string
	PropertyId    string
	BuyerId       string
}

func NewPendingPropertyTransaction(transactionId string, propertyId string, buyerId string, price int) *PendingPropertyTransaction {
	communication.FireEvent(communication.PROPERTY_TRANSFER_CREATED, NewPropertyTransferCreatedEvent(transactionId, "Bank", buyerId, price))
	return &PendingPropertyTransaction{TransactionId: transactionId, PropertyId: propertyId, BuyerId: buyerId}
}
