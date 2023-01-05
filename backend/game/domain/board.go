package domain

type Board struct {
	properties     []PropertyField
	eventFields    []EventField
	standardFields []StandardField
	prison         *Prison
}

func NewBoard(properties []PropertyField, eventFields []EventField, standardFields []StandardField, prison *Prison) *Board {
	return &Board{properties: properties, eventFields: eventFields, standardFields: standardFields, prison: prison}
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

func (board *Board) goToPrison(playerId string) {
	board.prison.addInmate(playerId)
}
