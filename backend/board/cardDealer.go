package board

import (
	"fmt"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
)

var currentCard *boardDomain.Card

func DrawCard() *boardDomain.Card {

	fmt.Println("Drawing a card..")
	//make random drawing
	card := &cardStack[0]
	currentCard = card
	eventing.FireEvent(eventing.CARD_DREW, boardDomain.NewCardDrewEvent(card.Title, card.Text))

	return card
}

//make more events
var cardStack = []boardDomain.Card{
	*boardDomain.NewCard("You inherited", "Your mentionned in the testament of your aunt. You receive 100 $.", func(player *boardDomain.Player) {
		fmt.Println("You inherited event...")
		eventing.FireEvent(eventing.PAYMENT, boardDomain.NewCreditAddedEvent(player.AccountId, 100))
	}),
}

func ConfirmCard() {

	currentCard.TriggerAction()
	currentCard = nil
}
