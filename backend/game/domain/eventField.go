package domain

type EventField struct {
	name      string
	index     int
	eventType EventType
}

type EventType string

const (
	DRAW_CARD   = "DRAW_CARD"
	GOTO_PRISON = "GOTO_PRISON"
	PAY_TAX     = "PAY_TAX"
)

func newEventField(name string, index int, eventType EventType) *EventField {
	return &EventField{name: name, index: index, eventType: eventType}
}

func (eventField EventField) Type() EventType {
	return eventField.eventType
}

func (eventField EventField) Name() string {
	return eventField.name
}

func (eventField EventField) Index() int {
	return eventField.index
}

var defaultEventFields = []EventField{
	*newEventField("Ereignisfeld 1", 4, DRAW_CARD),
	*newEventField("Ereignisfeld 2", 6, DRAW_CARD),
	*newEventField("Einkommenssteuer", 11, PAY_TAX),
	*newEventField("Gehe ins Gefaengnis", 13, GOTO_PRISON),
	*newEventField("Ereignisfeld 4", 15, DRAW_CARD),
}
