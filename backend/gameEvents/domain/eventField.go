package gameEventsDomain

type EventField struct {
	name  string
	id    string
	event func(playerId string)
}

func NewEventField(name string, id string, event func(playerId string)) *EventField {
	return &EventField{name: name, id: id, event: event}
}

func (eventField EventField) Name() string {
	return eventField.name
}

func (eventField EventField) Id() string {
	return eventField.id
}

func (eventField EventField) OnPlayerEnter(playerId string) {

	eventField.event(playerId)
}
