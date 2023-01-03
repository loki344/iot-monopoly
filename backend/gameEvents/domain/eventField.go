package gameEventsDomain

type EventField struct {
	name      string
	id        string
	eventType EventType
}

type EventType string

const (
	DRAW_CARD   = "DRAW_CARD"
	GOTO_PRISON = "GOTO_PRISON"
	PAY_TAX     = "PAY_TAX"
)

func NewEventField(name string, id string, eventType EventType) *EventField {
	return &EventField{name: name, id: id, eventType: eventType}
}

func (eventField EventField) Type() EventType {
	return eventField.eventType
}

func (eventField EventField) Name() string {
	return eventField.name
}

func (eventField EventField) Id() string {
	return eventField.id
}
