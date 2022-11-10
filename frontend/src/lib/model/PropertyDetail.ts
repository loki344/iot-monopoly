export type Property = {
	Id: String;
	Name: String;
	FinancialDetails: FinancialDetails;
	OwnerId: String;
	Upgrades?: PropertyUpgrades;
};

export enum PropertyUpgrades {
	ONE_HOUSE = 'oneHouse',
	TWO_HOUSES = 'twoHouses',
	THREE_HOUSES = 'threeHouses',
	FOUR_HOUSES = 'fourHouses',
	HOTEL = 'hotel'
}

export type FinancialDetails = {
	PropertyPrice: number;
	HousePrice: number;
	HotelPrice: number;
	Revenue: Revenue;
};

export type Revenue = {
	Normal: number;
	OneHouse: number;
	TwoHouses: number;
	ThreeHouses: number;
	FourHouses: number;
	Hotel: number;
};
