package repository

import (
	"iot-monopoly/eventing"
	domain "iot-monopoly/gameEvents/domain"
)

//make more events
var cardStack = domain.CardStack{Cards: []domain.Card{
	*domain.NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(playerId string) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED, domain.NewGameEventWithPayout(playerId, 100))
	}),
	*domain.NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(playerId string) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, domain.NewGameEventWithFee("Bank", playerId, 200))
	}),
}}

var currentCard *domain.Card

func GetNextCard(playerId string) {
	currentCard = cardStack.GetNextCard(playerId)
}

func GetCurrentCard() *domain.Card {
	return currentCard
}

func DeleteCard(card *domain.Card) {

	if card == currentCard {
		currentCard = nil
	}
}
