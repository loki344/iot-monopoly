package boardDomain

import (
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/eventing"
	financeDomain "iot-monopoly/finance/domain"
)

type Field interface {
	OnPlayerEnter(player *Player)
}

//TODO consolidate Name to super class?
type PropertyField struct {
	PropertyPrice uint64
	HousePrice    uint64
	HotelPrice    uint64
	Normal        uint64
	OneHouse      uint64
	TwoHouses     uint64
	ThreeHouses   uint64
	FourHouses    uint64
	Hotel         uint64
	Name          string
	Id            string
	OwnerId       string
	Upgrades      PropertyUpgrade
}

func NewPropertyField(name string, id string, financialDetails *FinancialDetails) *PropertyField {
	return &PropertyField{Name: name, Id: id, PropertyPrice: financialDetails.PropertyPrice, HousePrice: financialDetails.HousePrice, HotelPrice: financialDetails.HotelPrice,
		Normal: financialDetails.Revenue.Normal, OneHouse: financialDetails.Revenue.OneHouse, TwoHouses: financialDetails.Revenue.TwoHouses, ThreeHouses: financialDetails.Revenue.ThreeHouses,
		FourHouses: financialDetails.Revenue.FourHouses, Hotel: financialDetails.Revenue.Hotel}
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
	Normal      uint64
	OneHouse    uint64
	TwoHouses   uint64
	ThreeHouses uint64
	FourHouses  uint64
	Hotel       uint64
}

type FinancialDetails struct {
	PropertyPrice uint64
	HousePrice    uint64
	HotelPrice    uint64
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

func (propertyField PropertyField) GetPriceToPay() uint64 {
	switch propertyField.Upgrades {
	case ONE_HOUSE:
		return propertyField.OneHouse
	case TWO_HOUSES:
		return propertyField.TwoHouses
	case THREE_HOUSES:
		return propertyField.ThreeHouses
	case FOUR_HOUSES:
		return propertyField.FourHouses
	case HOTEL:
		return propertyField.Hotel
	default:
		return propertyField.Normal
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
		eventing.FireEvent(eventing.PROPERTY_BUY_QUESTION, NewPropertyBuyQuestion(player.Id, propertyField.Id))
	} else if propertyField.OwnerId == player.Id {
		fmt.Println("player owns the property..")
	} else {
		//TODO maybe it's cleaner to not fire the event here, fire it via service from finance domain?
		fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", propertyField.OwnerId, player.Id, propertyField.GetPriceToPay())
		eventing.FireEvent(eventing.TRANSACTION_REQUESTED, financeDomain.NewTransactionRequest(uuid.NewString(), propertyField.OwnerId, player.Id, propertyField.GetPriceToPay()))
	}
}
