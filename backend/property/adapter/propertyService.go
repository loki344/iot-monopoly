package adapter

import (
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/property/adapter/repository"
	domain "iot-monopoly/property/domain"
)

func BuyProperty(propertyId string, buyerId string) string {

	property := getPropertyById(propertyId)
	transactionId := uuid.NewString()
	repository.CreatePendingPropertyTransaction(domain.NewPendingPropertyTransaction(transactionId, propertyId, buyerId, property.FinancialDetails.PropertyPrice))

	return transactionId
}

func transferOwnerShip(transactionId string) {

	pendingTransfer := repository.GetPendingTransfer()

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.Id() == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyId(), pendingTransfer.BuyerId())
		property := getPropertyById(pendingTransfer.PropertyId())
		property.OwnerId = pendingTransfer.BuyerId()
		repository.DeletePendingPropertyTransaction(pendingTransfer)
	}
}

func getPropertyById(fieldId string) *domain.PropertyField {

	return repository.FindPropertyById(fieldId)
}
