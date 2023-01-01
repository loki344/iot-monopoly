package property

import (
	"fmt"
	"github.com/google/uuid"
	propertyDomain "iot-monopoly/property/domain"
)

var properties []*propertyDomain.PropertyField
var pendingTransfer *propertyDomain.PendingPropertyTransaction

var tempFinancialDetails = &propertyDomain.FinancialDetails{100, 100, 100, propertyDomain.Revenue{100, 200, 300, 400, 500, 800}}

//var defaultBasicFields = []propertyDomain.BasicField{
//	{propertyDomain.BaseFieldInformation{id: "1", name: "Start"}},
//	{propertyDomain.BaseFieldInformation{id: "5", name: "Gefaengnis"}},
//	{propertyDomain.BaseFieldInformation{id: "9", name: "Frei parken"}},
//}

var defaultProperties = []*propertyDomain.PropertyField{
	propertyDomain.NewPropertyField("Property purple 1", "2", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property purple 2", "3", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property orange 1", "7", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property orange 2", "8", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property green 1", "10", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property green 2", "12", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property blue 1", "14", tempFinancialDetails),
	propertyDomain.NewPropertyField("Property blue 2", "16", tempFinancialDetails),
}

func initFields() {
	properties = nil
	properties = defaultProperties
}

func BuyProperty(propertyId string, buyerId string) string {

	property := getPropertyById(propertyId)
	transactionId := uuid.NewString()
	pendingTransfer = propertyDomain.NewPendingPropertyTransaction(transactionId, propertyId, buyerId, property.FinancialDetails().PropertyPrice)

	return transactionId
}

func transferOwnerShip(transactionId string) {

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.Id() == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyId(), pendingTransfer.BuyerId())
		property := getPropertyById(pendingTransfer.PropertyId())
		property.SetOwnerId(pendingTransfer.BuyerId())
		pendingTransfer = nil
	}
}

func getPropertyById(fieldId string) *propertyDomain.PropertyField {

	for i := 0; i < len(properties); i++ {
		if properties[i].Id() == fieldId {
			return properties[i]
		}
	}
	panic("Field not found")
}
