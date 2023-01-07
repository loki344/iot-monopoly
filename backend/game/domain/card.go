package domain

import (
	"iot-monopoly/eventing"
)

type Card struct {
	title  string
	text   string
	action func(player *Player)
	player *Player
}

func (card *Card) SetPlayer(player *Player) {
	card.player = player
}

func (card Card) Title() string {
	return card.title
}

func (card Card) Text() string {
	return card.text
}

func (card Card) TriggerAction() {
	card.action(card.player)
}

func newCard(title string, text string, action func(player *Player)) *Card {

	return &Card{title: title, text: text, action: action}
}

var defaultCardStack = []Card{
	*newCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(player *Player) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED, newGameEventWithPayoutAcceptedEvent(player.Account().Id(), 100))
	}),
	*newCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(player *Player) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, newGameEventWithFeeAcceptedEvent("Bank", player.Account().Id(), 200))
	}),
	*newCard("Escape from prison", "Keep this card and escape from prison next time", func(player *Player) {
		player.IncreasePrisonCardCount()
	}),
}
