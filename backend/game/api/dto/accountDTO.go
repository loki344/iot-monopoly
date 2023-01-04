package dto

import "iot-monopoly/game/domain"

type AccountDTO struct {
	Id      string `json:"id"`
	Balance int    `json:"balance"`
}

func AccountDTOFromAccount(account *domain.Account) AccountDTO {
	return AccountDTO{Id: account.Id(), Balance: account.Balance()}
}
