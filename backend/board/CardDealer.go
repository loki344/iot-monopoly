package board

import (
	"fmt"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
)

func DrawCard() *boardDomain.CardEvent {

	fmt.Println("Drawing a card..")
	//make random drawing
	return &CardEvents[0]

}

//make more events
var CardEvents = []boardDomain.CardEvent{
	*boardDomain.NewCardEvent("You inherited", "Your mentionned in the testament of your aunt. You receive 100 $.", func(player *boardDomain.Player) {
		fmt.Println("You inherited event...")
		eventing.FireEvent(eventing.CREDIT, boardDomain.NewCreditAddedEvent(player.AccountId, 100))
	}),
}
