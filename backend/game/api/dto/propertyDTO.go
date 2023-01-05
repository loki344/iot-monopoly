package dto

import "iot-monopoly/game/domain"

type PropertyDTO struct {
	Name             string              `json:"name"`
	Index            int                 `json:"index"`
	FinancialDetails FinancialDetailsDTO `json:"financialDetails"`
	OwnerId          string              `json:"ownerId"`
	Upgrades         string              `json:"upgrades"`
}

type RevenueDTO struct {
	Normal      int `json:"normal"`
	OneHouse    int `json:"oneHouse"`
	TwoHouses   int `json:"twoHouses"`
	ThreeHouses int `json:"threeHouses"`
	FourHouses  int `json:"fourHouses"`
	Hotel       int `json:"hotel"`
}

type FinancialDetailsDTO struct {
	PropertyPrice int        `json:"propertyPrice"`
	HousePrice    int        `json:"housePrice"`
	HotelPrice    int        `json:"hotelPrice"`
	Revenue       RevenueDTO `json:"revenue"`
}

func PropertyDTOFromProperty(property *domain.PropertyField) PropertyDTO {
	return PropertyDTO{
		Name: property.Name(),
		FinancialDetails: FinancialDetailsDTO{
			PropertyPrice: property.FinancialDetails().PropertyPrice(),
			HousePrice:    property.FinancialDetails().HousePrice(),
			HotelPrice:    property.FinancialDetails().HotelPrice(),
			Revenue: RevenueDTO{
				Normal:      property.FinancialDetails().Revenue().Normal(),
				OneHouse:    property.FinancialDetails().Revenue().OneHouse(),
				TwoHouses:   property.FinancialDetails().Revenue().TwoHouses(),
				ThreeHouses: property.FinancialDetails().Revenue().ThreeHouses(),
				FourHouses:  property.FinancialDetails().Revenue().FourHouses(),
				Hotel:       property.FinancialDetails().Revenue().Hotel(),
			},
		},
		Upgrades: string(property.Upgrades()),
	}
}
