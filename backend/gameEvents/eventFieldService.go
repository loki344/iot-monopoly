package gameEvents

import (
	"fmt"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
	"iot-monopoly/player"
)

var eventFields []*gameEventsDomain.EventField

var defaultEventFields = []*gameEventsDomain.EventField{
	gameEventsDomain.NewEventField("Ereignisfeld 1", "4", func(playerId string) {
		DrawCard(playerId)
	}),
	gameEventsDomain.NewEventField("Ereignisfeld 2", "6", func(playerId string) {
		DrawCard(playerId)
	}),
	gameEventsDomain.NewEventField("Ereignisfeld 3", "11", func(playerId string) {
		DrawCard(playerId)
	}),
	gameEventsDomain.NewEventField("Gehe ins Gefaengnis", "13", func(playerId string) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		player.MovePlayer(playerId, 4)
	}),
	gameEventsDomain.NewEventField("Ereignisfeld 4", "15", func(playerId string) {
		DrawCard(playerId)
	}),
}

func initFields() {
	eventFields = nil
	eventFields = defaultEventFields
}

func GetFieldById(fieldId string) *gameEventsDomain.EventField {

	for i := 0; i < len(eventFields); i++ {
		if eventFields[i].Id() == fieldId {
			return eventFields[i]
		}
	}
	return nil
}

var currentCard *gameEventsDomain.Card

func DrawCard(playerId string) {

	fmt.Println("Drawing a card..")
	currentCard = gameEventsDomain.GetNextCard(playerId)
}

func ConfirmCard() {

	currentCard.TriggerAction()
	currentCard = nil
}
