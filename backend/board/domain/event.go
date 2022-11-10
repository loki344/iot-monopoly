package boardDomain

import eventingDomain "iot-monopoly/eventing/domain"

type PropertyBuyQuestion struct {
	PropertyId string
	PlayerId   string
}

func NewPropertyBuyQuestion(playerId string, propertyId string) *PropertyBuyQuestion {
	return &PropertyBuyQuestion{propertyId, playerId}
}

type LapFinishedEvent struct {
	*eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(LapFinishedEvent{}), playerId}
}
