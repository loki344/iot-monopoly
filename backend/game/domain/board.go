package domain

type Board struct {
	properties  []PropertyField
	eventFields []EventField
}

func NewBoard(properties []PropertyField, eventFields []EventField) *Board {
	return &Board{properties: properties, eventFields: eventFields}
}

func (board *Board) GetEventFieldByIndex(index int) *EventField {

	for i := range board.eventFields {
		if board.eventFields[i].index == index {
			return &board.eventFields[i]
		}
	}
	return nil
}

func (board *Board) GetPropertyByIndex(index int) *PropertyField {

	for i := range board.properties {
		if board.properties[i].index == index {
			return &board.properties[i]
		}
	}
	return nil
}
