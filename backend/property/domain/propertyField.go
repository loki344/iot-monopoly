package propertyDomain

import (
	"fmt"
	"iot-monopoly/communication"
)

type PropertyField struct {
	name             string
	id               string
	financialDetails *FinancialDetails
	ownerId          string
	upgrades         PropertyUpgrade
}

func (propertyField *PropertyField) SetOwnerId(ownerId string) {
	propertyField.ownerId = ownerId
}

func (propertyField PropertyField) Name() string {
	return propertyField.name
}

func (propertyField PropertyField) Id() string {
	return propertyField.id
}

func (propertyField PropertyField) FinancialDetails() *FinancialDetails {
	return propertyField.financialDetails
}

func (propertyField PropertyField) OwnerId() string {
	return propertyField.ownerId
}

func (propertyField PropertyField) Upgrades() PropertyUpgrade {
	return propertyField.upgrades
}

func NewPropertyField(name string, id string, financialDetails *FinancialDetails) *PropertyField {
	return &PropertyField{name: name, id: id, financialDetails: financialDetails}
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

func (propertyField PropertyField) GetPropertyFee() int {
	switch propertyField.upgrades {
	case ONE_HOUSE:
		return propertyField.financialDetails.Revenue.OneHouse
	case TWO_HOUSES:
		return propertyField.financialDetails.Revenue.TwoHouses
	case THREE_HOUSES:
		return propertyField.financialDetails.Revenue.ThreeHouses
	case FOUR_HOUSES:
		return propertyField.financialDetails.Revenue.FourHouses
	case HOTEL:
		return propertyField.financialDetails.Revenue.Hotel
	default:
		return propertyField.financialDetails.Revenue.Normal
	}
}

func (propertyField PropertyField) OnPlayerEnter(playerId string) {

	if propertyField.ownerId == "" {
		fmt.Println("property has no owner, is buyable")
		communication.FireEvent(communication.PLAYER_ON_UNOWNED_FIELD, NewPlayerOnUnownedFieldEvent(playerId, propertyField))
	} else if propertyField.ownerId != playerId {
		fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", propertyField.ownerId, playerId, propertyField.GetPropertyFee())
		communication.FireEvent(communication.PLAYER_ON_OWNED_FIELD, NewPlayerOnOwnedFieldEvent(playerId, propertyField))
	}
}
