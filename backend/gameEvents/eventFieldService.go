package gameEvents

import (
	"fmt"
	"iot-monopoly/board"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
)

var eventFields []*gameEventsDomain.EventField

var defaultEventFields = []*gameEventsDomain.EventField{
	{Name: "4", Id: "Ereignisfeld 1", Event: func(playerId string) {
		DrawCard(playerId)
	}},
	{Name: "6", Id: "Ereignisfeld 2", Event: func(playerId string) {
		DrawCard(playerId)
	}},
	{Name: "11", Id: "Ereignisfeld 3", Event: func(playerId string) {
		DrawCard(playerId)
	}},
	{Name: "13", Id: "Gehe ins Gefaegnis", Event: func(playerId string) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		board.MovePlayer(playerId, 4)
	}},
	{Name: "15", Id: "Ereignisfeld 4", Event: func(playerId string) {
		DrawCard(playerId)
	}},
}

func initFields() {
	eventFields = nil
	eventFields = defaultEventFields
}

func GetFieldById(fieldId string) *gameEventsDomain.EventField {

	for i := 0; i < len(eventFields); i++ {
		if eventFields[i].Id == fieldId {
			return eventFields[i]
		}
	}
	return nil
}
