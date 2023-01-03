package domain

import (
	"fmt"
	"iot-monopoly/eventing"
)

type PropertyField struct {
	Name             string            `json:"name"`
	Id               string            `json:"id"`
	FinancialDetails *FinancialDetails `json:"financialDetails"`
	OwnerId          string            `json:"ownerId"`
	Upgrades         PropertyUpgrade   `json:"upgrades"`
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

func (propertyField PropertyField) OnPlayerEnter(playerId string) {

	if propertyField.OwnerId == "" {
		fmt.Println("property has no owner, is buyable")
		eventing.FireEvent(eventing.PLAYER_ON_UNOWNED_FIELD, NewPlayerOnUnownedFieldEvent(playerId, propertyField))
	} else if propertyField.OwnerId != playerId {
		fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", propertyField.OwnerId, playerId, propertyField.GetPropertyFee())
		eventing.FireEvent(eventing.PLAYER_ON_OWNED_FIELD, NewPlayerOnOwnedFieldEvent(playerId, propertyField))
	}
}
