package gameEventsDomain

import (
	"iot-monopoly/eventing"
	"math/rand"
	"time"
)

type CardStack struct {
	Cards []Card
}

func (cardStack CardStack) GetNextCard(playerId string) *Card {

	rand.Seed(time.Now().UnixNano())
	card := &cardStack.Cards[rand.Intn(len(cardStack.Cards))]
	card.playerId = playerId

	eventing.FireEvent(eventing.CARD_DREW, NewCardDrewEvent(card.title, card.text))

	return card
}
