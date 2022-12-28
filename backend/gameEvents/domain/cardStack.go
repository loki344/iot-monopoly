package gameEventsDomain

import (
	"iot-monopoly/communication"
	"math/rand"
	"time"
)

//make more events
var CardStack = []Card{
	*NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(playerId string) {
		communication.FireEvent(communication.CARD_WITH_PAYOUT_ACCEPTED, NewCardWithPayoutDrewEvent(playerId, 100))
	}),
	*NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(playerId string) {
		communication.FireEvent(communication.CARD_WITH_FEE_ACCEPTED, NewCardWithFeeDrewEvent("Bank", playerId, 200))
	}),
}

func GetNextCard(playerId string) *Card {

	rand.Seed(time.Now().UnixNano())
	card := &CardStack[rand.Intn(len(CardStack))]
	card.PlayerId = playerId

	communication.FireEvent(communication.CARD_DREW, NewCardDrewEvent(card.Title, card.Text))

	return card
}
