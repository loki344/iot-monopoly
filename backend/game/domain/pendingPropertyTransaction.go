package domain

type PendingPropertyTransfer struct {
	id            string
	propertyIndex int
	buyerId       string
}

func (p PendingPropertyTransfer) Id() string {
	return p.id
}

func (p PendingPropertyTransfer) PropertyIndex() int {
	return p.propertyIndex
}

func (p PendingPropertyTransfer) BuyerId() string {
	return p.buyerId
}

func NewPendingPropertyTransaction(transactionId string, propertyIndex int, buyerId string) *PendingPropertyTransfer {
	return &PendingPropertyTransfer{id: transactionId, propertyIndex: propertyIndex, buyerId: buyerId}
}
