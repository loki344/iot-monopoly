package boardDomain

type Player struct {
	Id        string `json:"id"`
	Position  int    `json:"position"`
	AccountId string `json:"accountId"`
}
