package boardDomain

type Player struct {
	Id string `json:"id"`
	// TODO maybe change position to field
	Position int `json:"position"`
	Balance  int `json:"balance"`
}
