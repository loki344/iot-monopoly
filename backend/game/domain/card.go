package domain

type Card struct {
	title  string
	text   string
	action func(player *Player)
	player Player
}

func (card *Card) Player() Player {
	return card.player
}

func (card *Card) SetPlayer(player Player) {
	card.player = player
}

func (card Card) Title() string {
	return card.title
}

func (card Card) Text() string {
	return card.text
}

func (card Card) TriggerAction() {
	card.action(&card.player)
}

func NewCard(title string, text string, action func(player *Player)) *Card {

	return &Card{title: title, text: text, action: action}
}
