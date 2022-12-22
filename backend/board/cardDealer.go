package board

import (
	"fmt"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"math/rand"
	"time"
)

var currentCard *boardDomain.Card

func DrawCard() *boardDomain.Card {

	fmt.Println("Drawing a card..")

	rand.Seed(time.Now().UnixNano())
	card := &cardStack[rand.Intn(len(cardStack))]
	currentCard = card
	eventing.FireEvent(eventing.CARD_DREW, boardDomain.NewCardDrewEvent(card.Title, card.Text))
	return card
}

//make more events
var cardStack = []boardDomain.Card{
	*boardDomain.NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(player *boardDomain.Player) {
		eventing.FireEvent(eventing.PAYOUT_REQUESTED, boardDomain.NewCreditAddedEvent(player.AccountId, 100))
	}),
	*boardDomain.NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(player *boardDomain.Player) {
		eventing.FireEvent(eventing.PLAYER_ON_OWNED_FIELD, boardDomain.NewTransactionRequest("Bank", player.Id, 200))
	}),
}

func ConfirmCard() {

	currentCard.TriggerAction()
	currentCard = nil
}
