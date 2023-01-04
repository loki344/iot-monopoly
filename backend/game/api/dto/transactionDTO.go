package dto

import "iot-monopoly/game/domain"

type TransactionDTO struct {
	Id          string `json:"id"`
	RecipientId string `json:"recipientId"`
	SenderId    string `json:"senderId"`
	Amount      int    `json:"amount"`
	Accepted    bool   `json:"accepted"`
}

func TransactionDTOFromTransaction(transaction *domain.Transaction) TransactionDTO {
	return TransactionDTO{
		Id:          transaction.Id(),
		RecipientId: transaction.RecipientId(),
		SenderId:    transaction.SenderId(),
		Amount:      transaction.Amount(),
		Accepted:    transaction.Accepted(),
	}
}
