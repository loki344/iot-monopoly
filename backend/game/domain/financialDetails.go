package domain

type FinancialDetails struct {
	propertyPrice int
	housePrice    int
	hotelPrice    int
	revenue       Revenue
}

func (f FinancialDetails) PropertyPrice() int {
	return f.propertyPrice
}

func (f FinancialDetails) HousePrice() int {
	return f.housePrice
}

func (f FinancialDetails) HotelPrice() int {
	return f.hotelPrice
}

func (f FinancialDetails) Revenue() Revenue {
	return f.revenue
}
