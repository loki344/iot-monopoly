package gameEventsDomain

type Card struct {
	title    string
	text     string
	action   func(playerId string)
	playerId string
}

func (card Card) Title() string {
	return card.title
}

func (card Card) Text() string {
	return card.text
}

func (card Card) PlayerId() string {
	return card.playerId
}

func (card Card) TriggerAction() {
	card.action(card.playerId)
}

func NewCard(title string, text string, action func(playerId string)) *Card {

	return &Card{title: title, text: text, action: action}
}
