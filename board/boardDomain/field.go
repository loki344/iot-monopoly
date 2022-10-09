package boardDomain

import "fmt"

type Field interface {
	OnPlayerEnter(player *Player)
}

//TODO consolidate Name to super class?
type PropertyField struct {
	Name  string
	Price int
	Owner *Player
}

type EventField struct {
	Name  string
	Event func(player *Player)
}

type BasicField struct {
	Name string
}

func (eventField EventField) OnPlayerEnter(player *Player) {

	eventField.Event(player)
}

func (_ BasicField) OnPlayerEnter(player *Player) {
	// do nothing
}

func (propertyField PropertyField) OnPlayerEnter(player *Player) {

	if propertyField.Owner == player {
		fmt.Println("player owns the property..")
	} else {
		fmt.Println("player does not own the property..")

	}
}
