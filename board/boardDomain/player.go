package boardDomain

type Player struct {
	Id string
	// TODO maybe change position to field
	Position int
	Balance  int
}

type LapFinishedEvent struct {
	PlayerId string
}
