package domain

type PropertyField struct {
	name             string
	index            int
	financialDetails FinancialDetails
	ownerId          string
	upgrades         PropertyUpgrade
}

func (propertyField PropertyField) Name() string {
	return propertyField.name
}

func (propertyField PropertyField) FinancialDetails() FinancialDetails {
	return propertyField.financialDetails
}

func (propertyField PropertyField) Upgrades() PropertyUpgrade {
	return propertyField.upgrades
}

func NewPropertyField(name string, index int, financialDetails FinancialDetails) *PropertyField {
	return &PropertyField{name: name, index: index, financialDetails: financialDetails}
}

type PropertyUpgrade string

const (
	ONE_HOUSE    PropertyUpgrade = "oneHouse"
	TWO_HOUSES   PropertyUpgrade = "twoHouses"
	THREE_HOUSES PropertyUpgrade = "threeHouses"
	FOUR_HOUSES  PropertyUpgrade = "fourHouses"
	HOTEL        PropertyUpgrade = "hotel"
)

func (propertyField PropertyField) GetPropertyFee() int {
	switch propertyField.upgrades {
	case ONE_HOUSE:
		return propertyField.financialDetails.revenue.oneHouse
	case TWO_HOUSES:
		return propertyField.financialDetails.revenue.twoHouses
	case THREE_HOUSES:
		return propertyField.financialDetails.revenue.threeHouses
	case FOUR_HOUSES:
		return propertyField.financialDetails.revenue.fourHouses
	case HOTEL:
		return propertyField.financialDetails.revenue.hotel
	default:
		return propertyField.financialDetails.revenue.normal
	}
}

func (propertyField PropertyField) GetPrice() int {
	return propertyField.financialDetails.propertyPrice
}
