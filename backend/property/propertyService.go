package property

import (
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/communication"
	"iot-monopoly/property/domain"
)

var properties []*propertyDomain.PropertyField
var pendingTransfer *propertyDomain.PendingPropertyTransaction

var tempFinancialDetails = &propertyDomain.FinancialDetails{100, 100, 100, propertyDomain.Revenue{100, 200, 300, 400, 500, 800}}

//var defaultBasicFields = []propertyDomain.BasicField{
//	{propertyDomain.BaseFieldInformation{Id: "1", Name: "Start"}},
//	{propertyDomain.BaseFieldInformation{Id: "5", Name: "Gefaengnis"}},
//	{propertyDomain.BaseFieldInformation{Id: "9", Name: "Frei parken"}},
//}

var defaultProperties = []*propertyDomain.PropertyField{
	propertyDomain.NewPropertyField("2", "Property purple 1", tempFinancialDetails),
	propertyDomain.NewPropertyField("3", "Property purple 2", tempFinancialDetails),
	propertyDomain.NewPropertyField("7", "Property orange 1", tempFinancialDetails),
	propertyDomain.NewPropertyField("8", "Property orange 2", tempFinancialDetails),
	propertyDomain.NewPropertyField("10", "Property green 1", tempFinancialDetails),
	propertyDomain.NewPropertyField("12", "Property green 2", tempFinancialDetails),
	propertyDomain.NewPropertyField("14", "Property blue 1", tempFinancialDetails),
	propertyDomain.NewPropertyField("16", "Property blue 2", tempFinancialDetails),
}

func initFields() {
	properties = nil
	properties = defaultProperties
}

func BuyProperty(propertyId string, buyerId string) string {

	property := *GetPropertyById(propertyId)
	transactionId := uuid.NewString()
	pendingTransfer = &propertyDomain.PendingPropertyTransaction{TransactionId: transactionId, PropertyId: propertyId, BuyerId: buyerId}
	communication.FireEvent(communication.PROPERTY_TRANSFER_CREATED, propertyDomain.NewPropertyTransferCreatedEvent(transactionId, "Bank", buyerId, property.GetPropertyFee()))

	return transactionId
}

func transferOwnerShip(transactionId string) {

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.TransactionId == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyId, pendingTransfer.BuyerId)
		property := GetPropertyById(pendingTransfer.PropertyId)
		property.OwnerId = pendingTransfer.BuyerId
		pendingTransfer = nil
	}
}

func GetPropertyById(fieldId string) *propertyDomain.PropertyField {

	for i := 0; i < len(properties); i++ {
		if properties[i].Id == fieldId {
			return properties[i]
		}
	}
	panic("Field not found")
}
