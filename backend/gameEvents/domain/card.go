package gameEventsDomain

import (
	eventingDomain "iot-monopoly/communication/domain"
)

type Card struct {
	Title    string
	Text     string
	Action   func(playerId string)
	PlayerId string
}

func (card Card) TriggerAction() {
	card.Action(card.PlayerId)
}

type CardDTO struct {
	eventingDomain.BaseEvent
	Title string
	Text  string
}

func NewCard(title string, text string, action func(playerId string)) *Card {

	return &Card{Title: title, Text: text, Action: action}
}
