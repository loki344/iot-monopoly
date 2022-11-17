package boardDomain

import (
	"fmt"
	"iot-monopoly/eventing"
)

type Field interface {
	OnPlayerEnter(player *Player)
	GetId() string
	GetPriceToPay() int
}

//TODO consolidate Name to super class?
type PropertyField struct {
	Name             string
	Id               string
	FinancialDetails *FinancialDetails
	OwnerId          string
	Upgrades         PropertyUpgrade
}

func (propertyField PropertyField) GetId() string {
	return propertyField.Id
}

func NewPropertyField(name string, id string, financialDetails *FinancialDetails) *PropertyField {
	return &PropertyField{Name: name, Id: id, FinancialDetails: financialDetails}
}

type PropertyUpgrade string

const (
	ONE_HOUSE    PropertyUpgrade = "oneHouse"
	TWO_HOUSES   PropertyUpgrade = "twoHouses"
	THREE_HOUSES PropertyUpgrade = "threeHouses"
	FOUR_HOUSES  PropertyUpgrade = "fourHouses"
	HOTEL        PropertyUpgrade = "hotel"
)

type Revenue struct {
	Normal      int
	OneHouse    int
	TwoHouses   int
	ThreeHouses int
	FourHouses  int
	Hotel       int
}

type FinancialDetails struct {
	PropertyPrice int
	HousePrice    int
	HotelPrice    int
	Revenue       Revenue
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

func (field BasicField) GetId() string {
	return field.Id
}

func (eventField EventField) GetId() string {
	return eventField.Id
}
func (field BasicField) GetPriceToPay() int {
	panic("BasicField is not buyable and has no price")
}

func (eventField EventField) GetPriceToPay() int {
	panic("EventField is not buyable and has no price")
}

func (propertyField PropertyField) GetPriceToPay() int {
	switch propertyField.Upgrades {
	case ONE_HOUSE:
		return propertyField.FinancialDetails.Revenue.OneHouse
	case TWO_HOUSES:
		return propertyField.FinancialDetails.Revenue.TwoHouses
	case THREE_HOUSES:
		return propertyField.FinancialDetails.Revenue.ThreeHouses
	case FOUR_HOUSES:
		return propertyField.FinancialDetails.Revenue.FourHouses
	case HOTEL:
		return propertyField.FinancialDetails.Revenue.Hotel
	default:
		return propertyField.FinancialDetails.Revenue.Normal
	}
}

func (eventField EventField) OnPlayerEnter(player *Player) {

	eventField.Event(player)
}

func (_ BasicField) OnPlayerEnter(_ *Player) {
	// do nothing
}

func (propertyField PropertyField) OnPlayerEnter(player *Player) {

	if propertyField.OwnerId == "" {
		fmt.Println("property has no owner, is buyable")
		eventing.FireEvent(eventing.PROPERTY_BUY_QUESTION, NewPropertyBuyQuestion(player.Id, propertyField))
	} else if propertyField.OwnerId == player.Id {
		fmt.Println("player owns the property..")
	} else {
		fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", propertyField.OwnerId, player.Id, propertyField.GetPriceToPay())
		eventing.FireEvent(eventing.PROPERTY_FEE, NewPropertyFeeRequest(propertyField.OwnerId, player.Id, propertyField.GetPriceToPay()))
	}
}
