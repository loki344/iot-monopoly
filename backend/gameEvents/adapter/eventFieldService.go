package gameEventsAdapter

import (
	"fmt"
	domain "iot-monopoly/gameEvents/domain"
)

var eventFields []*domain.EventField

var defaultEventFields = []*domain.EventField{
	domain.NewEventField("Ereignisfeld 1", "4", func(playerId string) {
		DrawCard(playerId)
	}),
	domain.NewEventField("Ereignisfeld 2", "6", func(playerId string) {
		DrawCard(playerId)
	}),
	domain.NewEventField("Ereignisfeld 3", "11", func(playerId string) {
		DrawCard(playerId)
	}),
	domain.NewEventField("Gehe ins Gefaengnis", "13", func(playerId string) {
		fmt.Println("Player has to go to prison")
		// TODO implement
	}),
	domain.NewEventField("Ereignisfeld 4", "15", func(playerId string) {
		DrawCard(playerId)
	}),
}

func initFields() {
	eventFields = nil
	eventFields = defaultEventFields
}

func GetFieldById(fieldId string) *domain.EventField {

	for i := 0; i < len(eventFields); i++ {
		if eventFields[i].Id() == fieldId {
			return eventFields[i]
		}
	}
	return nil
}

var currentCard *domain.Card

func DrawCard(playerId string) {

	fmt.Println("Drawing a card..")
	currentCard = domain.GetNextCard(playerId)
}

func ConfirmCard() {

	currentCard.TriggerAction()
	currentCard = nil
}
