package financeDomain

type Account struct {
	Id       string `json:"id"`
	PlayerId string `json:"playerId"`
	Balance  int    `json:"balance"`
}
