package domain

import eventingDomain "iot-monopoly/eventing/domain"

type CardDrewEvent struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCardDrewEvent(title string, text string) *CardDrewEvent {
	return &CardDrewEvent{BaseEvent: eventingDomain.EventType(&CardDrewEvent{}), Title: title, Text: text}
}
