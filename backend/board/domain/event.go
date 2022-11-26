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

type GameStartedEvent struct {
	eventingDomain.BaseEvent
	PlayerCount int
}

func NewGameStartedEvent(playerCount int) *GameStartedEvent {
	return &GameStartedEvent{eventingDomain.EventType(&GameStartedEvent{}), playerCount}
}

type CreditAddedEvent struct {
	eventingDomain.BaseEvent
	RecipientAccountId string
	Amount             int
}

func NewCreditAddedEvent(recipientAccountId string, amount int) *CreditAddedEvent {

	return &CreditAddedEvent{BaseEvent: eventingDomain.EventType(&CreditAddedEvent{}), RecipientAccountId: recipientAccountId, Amount: amount}
}

type CardEvent struct {
	eventingDomain.BaseEvent
	Title  string
	Text   string
	Event  func(player *Player)
	Player *Player
}

type CardEventDTO struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCardEvent(title string, text string, event func(player *Player)) *CardEvent {

	return &CardEvent{BaseEvent: eventingDomain.EventType(&CardEvent{}), Title: title, Text: text, Event: event}
}
