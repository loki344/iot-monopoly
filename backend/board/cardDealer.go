package board

import (
	"fmt"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/communication"
	"math/rand"
	"time"
)

var currentCard *boardDomain.Card

func DrawCard() *boardDomain.Card {

	fmt.Println("Drawing a card..")

	rand.Seed(time.Now().UnixNano())
	card := &cardStack[rand.Intn(len(cardStack))]
	currentCard = card
	communication.FireEvent(communication.CARD_DREW, boardDomain.NewCardDrewEvent(card.Title, card.Text))
	return card
}

//make more events
var cardStack = []boardDomain.Card{
	*boardDomain.NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(playerId string) {
		communication.FireEvent(communication.CARD_WITH_PAYOUT_DREW, boardDomain.NewCardWithPayoutDrewEvent(playerId, 100))
	}),
	*boardDomain.NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(playerId string) {
		communication.FireEvent(communication.CARD_WITH_FEE_DREW, boardDomain.NewCardWithFeeDrewEvent("Bank", playerId, 200))
	}),
}

func ConfirmCard() {

	currentCard.TriggerAction()
	currentCard = nil
}
