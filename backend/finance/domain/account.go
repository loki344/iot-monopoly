package financeDomain

type Account struct {
	Id       string `json:"id"`
	PlayerId string `json:"accountId"`
	Balance  int    `json:"balance"`
}
