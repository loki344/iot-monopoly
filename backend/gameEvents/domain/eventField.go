package gameEventsDomain

type EventField struct {
	Name  string
	Id    string
	Event func(playerId string)
}

func (eventField EventField) OnPlayerEnter(playerId string) {

	eventField.Event(playerId)
}
