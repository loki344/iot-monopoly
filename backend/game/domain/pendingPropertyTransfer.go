package domain

type PendingPropertyTransfer struct {
	referenceTransactionId string
	propertyIndex          int
	buyerId                string
}

func (p PendingPropertyTransfer) ReferenceTransactionId() string {
	return p.referenceTransactionId
}

func (p PendingPropertyTransfer) PropertyIndex() int {
	return p.propertyIndex
}

func (p PendingPropertyTransfer) BuyerId() string {
	return p.buyerId
}

func newPendingPropertyTransaction(transactionId string, propertyIndex int, buyerId string) *PendingPropertyTransfer {
	return &PendingPropertyTransfer{referenceTransactionId: transactionId, propertyIndex: propertyIndex, buyerId: buyerId}
}
