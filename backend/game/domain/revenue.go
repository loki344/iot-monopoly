package domain

type Revenue struct {
	normal      int
	oneHouse    int
	twoHouses   int
	threeHouses int
	fourHouses  int
	hotel       int
}

func (r Revenue) Normal() int {
	return r.normal
}

func (r Revenue) OneHouse() int {
	return r.oneHouse
}

func (r Revenue) TwoHouses() int {
	return r.twoHouses
}

func (r Revenue) ThreeHouses() int {
	return r.threeHouses
}

func (r Revenue) FourHouses() int {
	return r.fourHouses
}

func (r Revenue) Hotel() int {
	return r.hotel
}
