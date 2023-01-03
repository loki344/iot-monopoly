package propertyAdapter

import (
	"fmt"
	"github.com/google/uuid"
	domain "iot-monopoly/property/domain"
)

var properties []*domain.PropertyField
var pendingTransfer *domain.PendingPropertyTransaction

var tempFinancialDetails = &domain.FinancialDetails{100, 100, 100, domain.Revenue{100, 200, 300, 400, 500, 800}}

//var defaultBasicFields = []domain.BasicField{
//	{domain.BaseFieldInformation{id: "1", name: "Start"}},
//	{domain.BaseFieldInformation{id: "5", name: "Gefaengnis"}},
//	{domain.BaseFieldInformation{id: "9", name: "Frei parken"}},
//}

var defaultProperties = []*domain.PropertyField{
	domain.NewPropertyField("Property purple 1", "2", tempFinancialDetails),
	domain.NewPropertyField("Property purple 2", "3", tempFinancialDetails),
	domain.NewPropertyField("Property orange 1", "7", tempFinancialDetails),
	domain.NewPropertyField("Property orange 2", "8", tempFinancialDetails),
	domain.NewPropertyField("Property green 1", "10", tempFinancialDetails),
	domain.NewPropertyField("Property green 2", "12", tempFinancialDetails),
	domain.NewPropertyField("Property blue 1", "14", tempFinancialDetails),
	domain.NewPropertyField("Property blue 2", "16", tempFinancialDetails),
}

func initFields() {
	properties = nil
	properties = defaultProperties
}

func BuyProperty(propertyId string, buyerId string) string {

	property := getPropertyById(propertyId)
	transactionId := uuid.NewString()
	pendingTransfer = domain.NewPendingPropertyTransaction(transactionId, propertyId, buyerId, property.FinancialDetails.PropertyPrice)

	return transactionId
}

func transferOwnerShip(transactionId string) {

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.Id() == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyId(), pendingTransfer.BuyerId())
		property := getPropertyById(pendingTransfer.PropertyId())
		property.OwnerId = pendingTransfer.BuyerId()
		pendingTransfer = nil
	}
}

func getPropertyById(fieldId string) *domain.PropertyField {

	for i := 0; i < len(properties); i++ {
		if properties[i].Id == fieldId {
			return properties[i]
		}
	}
	return nil
}
