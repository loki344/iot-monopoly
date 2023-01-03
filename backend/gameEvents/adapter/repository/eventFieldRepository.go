package repository

import (
	domain "iot-monopoly/gameEvents/domain"
)

var eventFields []*domain.EventField

var defaultEventFields = []*domain.EventField{
	domain.NewEventField("Ereignisfeld 1", "4", domain.DRAW_CARD),
	domain.NewEventField("Ereignisfeld 2", "6", domain.DRAW_CARD),
	domain.NewEventField("Einkommenssteuer", "11", domain.PAY_TAX),
	domain.NewEventField("Gehe ins Gefaengnis", "13", domain.GOTO_PRISON),
	domain.NewEventField("Ereignisfeld 4", "15", domain.DRAW_CARD),
}

func InitFields() {
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
