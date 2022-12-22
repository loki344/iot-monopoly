package propertyDomain

import (
	"fmt"
	"iot-monopoly/communication"
)

type Field interface {
	OnPlayerEnter(playerId string)
	GetId() string
}

type BaseFieldInformation struct {
	Name string
	Id   string
}

type PropertyField struct {
	BaseFieldInformation
	FinancialDetails *FinancialDetails
	OwnerId          string
	Upgrades         PropertyUpgrade
}

func (propertyField PropertyField) GetId() string {
	return propertyField.Id
}

func NewPropertyField(baseInformation BaseFieldInformation, financialDetails *FinancialDetails) *PropertyField {
	return &PropertyField{BaseFieldInformation: baseInformation, FinancialDetails: financialDetails}
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
	BaseFieldInformation
	Event func(playerId string)
}

type BasicField struct {
	BaseFieldInformation
}

func (field BasicField) GetId() string {
	return field.Id
}

func (eventField EventField) GetId() string {
	return eventField.Id
}

func (propertyField PropertyField) GetPropertyFee() int {
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

func (eventField EventField) OnPlayerEnter(playerId string) {

	eventField.Event(playerId)
}

func (_ BasicField) OnPlayerEnter(_ string) {
	// do nothing
}

func (propertyField PropertyField) OnPlayerEnter(playerId string) {

	if propertyField.OwnerId == "" {
		fmt.Println("property has no owner, is buyable")
		communication.FireEvent(communication.PLAYER_ON_UNOWNED_FIELD, NewPlayerOnUnownedFieldEvent(playerId, propertyField))
	} else if propertyField.OwnerId != playerId {
		fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", propertyField.OwnerId, playerId, propertyField.GetPropertyFee())
		communication.FireEvent(communication.PLAYER_ON_OWNED_FIELD, NewPlayerOnOwnedFieldEvent(playerId, propertyField))
	}
}
