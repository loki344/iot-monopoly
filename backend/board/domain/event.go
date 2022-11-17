package boardDomain

import eventingDomain "iot-monopoly/eventing/domain"

type PropertyBuyQuestion struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property PropertyField
}

func NewPropertyBuyQuestion(playerId string, property PropertyField) *PropertyBuyQuestion {
	return &PropertyBuyQuestion{eventingDomain.EventType(&PropertyBuyQuestion{}), playerId, property}
}

type PropertyFeeRequest struct {
	eventingDomain.BaseEvent
	OwnerId string
	GuestId string
	Price   int
}

func NewPropertyFeeRequest(ownerId string, guestId string, price int) *PropertyFeeRequest {
	return &PropertyFeeRequest{eventingDomain.EventType(&PropertyFeeRequest{}), ownerId, guestId, price}
}

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}
