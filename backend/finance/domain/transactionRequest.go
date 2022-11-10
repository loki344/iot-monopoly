package financeDomain

type TransactionRequest struct {
	RecipientId string
	SenderId    string
	Amount      int
}
