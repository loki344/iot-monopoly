package repository

import (
	"iot-monopoly/game/domain"
)

var pendingTransaction *domain.Transaction

func GetPendingTransaction() *domain.Transaction {
	return pendingTransaction
}

func DeleteTransaction(transaction *domain.Transaction) {
	if transaction.Id() == pendingTransaction.Id() {
		pendingTransaction = nil
	}
}

func CreateTransaction(transaction *domain.Transaction) {
	pendingTransaction = transaction
}
