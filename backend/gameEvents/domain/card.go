package gameEventsDomain

type Card struct {
	Title    string
	Text     string
	Action   func(playerId string)
	PlayerId string
}

func (card Card) TriggerAction() {
	card.Action(card.PlayerId)
}

func NewCard(title string, text string, action func(playerId string)) *Card {

	return &Card{Title: title, Text: text, Action: action}
}
