package gameEvents

import (
	"fmt"
	"iot-monopoly/communication"
	gameEventsDomain "iot-monopoly/gameEvents/domain"
	"math/rand"
	"time"
)

var currentCard *gameEventsDomain.Card

func DrawCard(playerId string) *gameEventsDomain.Card {

	fmt.Println("Drawing a card..")

	rand.Seed(time.Now().UnixNano())
	card := &cardStack[rand.Intn(len(cardStack))]
	card.PlayerId = playerId
	currentCard = card
	communication.FireEvent(communication.CARD_DREW, gameEventsDomain.NewCardDrewEvent(card.Title, card.Text))
	return card
}

//make more events
var cardStack = []gameEventsDomain.Card{
	*gameEventsDomain.NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(playerId string) {
		communication.FireEvent(communication.CARD_WITH_PAYOUT_DREW, gameEventsDomain.NewCardWithPayoutDrewEvent(playerId, 100))
	}),
	*gameEventsDomain.NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(playerId string) {
		communication.FireEvent(communication.CARD_WITH_FEE_DREW, gameEventsDomain.NewCardWithFeeDrewEvent("Bank", playerId, 200))
	}),
}

func ConfirmCard() {

	currentCard.TriggerAction()
	currentCard = nil
}
