package eventingDomain

import "github.com/google/uuid"

type LapFinishedEvent struct {
	PlayerId string
}

type SensorEvent struct {
	PlayerId   string
	FieldIndex int
}

type PropertyBuyQuestion struct {
	PlayerId   string
	PropertyId string
}

type TransactionRequested struct {
	Id          string
	RecipientId string
	SenderId    string
	Amount      int
}

func NewTransactionRequest(recipientId string, senderId string, amount int) TransactionRequested {

	if amount <= 0 {
		panic("amount has to be greater than 0")
	}
	id := uuid.New().String()

	return TransactionRequested{Id: id, RecipientId: recipientId, SenderId: senderId, Amount: amount}
}
