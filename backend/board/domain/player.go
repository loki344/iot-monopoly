package boardDomain

type Player struct {
	Id string `json:"id"`
	// TODO maybe change position to field
	Position uint64 `json:"position"`
	Balance  uint64 `json:"balance"`
}
