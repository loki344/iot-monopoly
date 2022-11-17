package boardDomain

import (
	"github.com/google/uuid"
	eventingDomain "iot-monopoly/eventing/domain"
)

type PropertyBuyQuestion struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func NewPropertyBuyQuestion(playerId string, property PropertyField) *PropertyBuyQuestion {
	return &PropertyBuyQuestion{eventingDomain.EventType(&PropertyBuyQuestion{}), playerId, property}
}

type TransactionRequest struct {
	eventingDomain.BaseEvent
	TransactionId string
	ReceiverId    string
	SenderId      string
	Price         int
}

func NewTransactionRequestWithId(id string, receiverId string, senderId string, price int) *TransactionRequest {
	return &TransactionRequest{eventingDomain.EventType(&TransactionRequest{}), id, receiverId, senderId, price}
}

func NewTransactionRequest(receiverId string, senderId string, price int) *TransactionRequest {
	return NewTransactionRequestWithId(uuid.NewString(), receiverId, senderId, price)
}

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}
