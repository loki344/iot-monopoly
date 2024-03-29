package domain

type PropertyField struct {
	name             string
	index            int
	financialDetails FinancialDetails
	ownerId          string
	upgrades         PropertyUpgrade
}

func (propertyField PropertyField) OwnerId() string {
	return propertyField.ownerId
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

func newPropertyField(name string, index int, financialDetails FinancialDetails) *PropertyField {
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

var defaultProperties = []PropertyField{
	*newPropertyField("Property purple 1", 2, defaultFinancialDetails),
	*newPropertyField("Property purple 2", 3, defaultFinancialDetails),
	*newPropertyField("Property orange 1", 7, defaultFinancialDetails),
	*newPropertyField("Property orange 2", 8, defaultFinancialDetails),
	*newPropertyField("Property green 1", 10, defaultFinancialDetails),
	*newPropertyField("Property green 2", 12, defaultFinancialDetails),
	*newPropertyField("Property blue 1", 14, defaultFinancialDetails),
	*newPropertyField("Property blue 2", 16, defaultFinancialDetails),
}
var defaultFinancialDetails = FinancialDetails{100, 100, 100, Revenue{100, 200, 300, 400, 500, 800}}
