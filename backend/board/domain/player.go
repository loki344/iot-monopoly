package boardDomain

type Player struct {
	Id string `json:"id"`
	// TODO maybe change position to field
	Position uint8  `json:"position"`
	Balance  uint32 `json:"balance"`
}
