package boardDomain

import (
	"fmt"
	"iot-monopoly/eventing"
)

type Field interface {
	OnPlayerEnter(player *Player)
}

//TODO consolidate Name to super class?
type PropertyField struct {
	Name    string
	Id      string
	Price   int
	OwnerId string
}

type EventField struct {
	Name  string
	Id    string
	Event func(player *Player)
}

type BasicField struct {
	Name string
	Id   string
}

func (eventField EventField) OnPlayerEnter(player *Player) {

	eventField.Event(player)
}

func (_ BasicField) OnPlayerEnter(player *Player) {
	// do nothing
}

func (propertyField PropertyField) OnPlayerEnter(player *Player) {

	if propertyField.OwnerId == "" {
		fmt.Println("property has no owner, is buyable")
		eventing.FireEvent(eventing.PROPERTY_BUY_QUESTION, eventing.PropertyBuyQuestion{PlayerId: player.Id, PropertyId: propertyField.Id})
	} else if propertyField.OwnerId == player.Id {
		fmt.Println("player owns the property..")
	} else {
		fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", propertyField.OwnerId, player.Id, propertyField.Price)
		eventing.FireEvent(eventing.TRANSACTION_REQUESTED, eventing.NewTransactionRequest(propertyField.OwnerId, player.Id, propertyField.Price))
	}
}
