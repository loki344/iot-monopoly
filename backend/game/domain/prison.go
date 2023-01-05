package domain

type Prison struct {
	index   int
	inmates []Inmate
}

type Inmate struct {
	playerId       string
	roundsInPrison int
}

func (p Prison) addInmate(playerId string) {
	p.inmates = append(p.inmates, Inmate{playerId: playerId})
}

func (p Prison) isInmate(playerId string) bool {
	for i := range p.inmates {
		if p.inmates[i].playerId == playerId {
			return true
		}
	}
	return false
}

func newPrison() *Prison {
	return &Prison{index: 5}
}
