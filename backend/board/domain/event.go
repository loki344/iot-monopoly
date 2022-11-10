package boardDomain

import eventingDomain "iot-monopoly/eventing/domain"

type PropertyBuyQuestion struct {
	eventingDomain.BaseEvent
	PlayerId string
	Property *PropertyField
}

func NewPropertyBuyQuestion(playerId string, property *PropertyField) *PropertyBuyQuestion {
	return &PropertyBuyQuestion{eventingDomain.EventType(&PropertyBuyQuestion{}), playerId, property}
}

type LapFinishedEvent struct {
	eventingDomain.BaseEvent
	PlayerId string
}

func NewLapFinishedEvent(playerId string) *LapFinishedEvent {
	return &LapFinishedEvent{eventingDomain.EventType(&LapFinishedEvent{}), playerId}
}
