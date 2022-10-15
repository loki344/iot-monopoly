package eventingDomain

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
	id          string
	recipientId string
	senderId    string
	amount      int
}

func (t TransactionRequested) Id() string {
	return t.id
}

func (t TransactionRequested) RecipientId() string {
	return t.recipientId
}

func (t TransactionRequested) SenderId() string {
	return t.senderId
}

func (t TransactionRequested) Amount() int {
	return t.amount
}

func NewTransactionRequest(id string, recipientId string, senderId string, amount int) TransactionRequested {

	if amount <= 0 {
		panic("amount has to be greater than 0")
	}
	return TransactionRequested{id: id, recipientId: recipientId, senderId: senderId, amount: amount}
}
