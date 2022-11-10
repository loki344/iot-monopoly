package boardDomain

import eventingDomain "iot-monopoly/eventing/domain"

type PropertyBuyQuestion struct {
	Property *PropertyField
	eventingDomain.BaseEvent
	PlayerId string
}

func NewPropertyBuyQuestion(playerId string, property *PropertyField) *PropertyBuyQuestion {
	return &PropertyBuyQuestion{property, eventingDomain.EventType(&PropertyBuyQuestion{}), playerId}
}

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}
