package property

import (
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/board"
	"iot-monopoly/communication"
	"iot-monopoly/property/domain"
)

var properties []propertyDomain.PropertyField
var pendingTransfer *propertyDomain.PendingPropertyTransaction
var basicFields []propertyDomain.BasicField
var eventFields []propertyDomain.EventField

var tempFinancialDetails = &propertyDomain.FinancialDetails{100, 100, 100, propertyDomain.Revenue{100, 200, 300, 400, 500, 800}}

var defaultEventFields = []propertyDomain.EventField{
	{propertyDomain.BaseFieldInformation{Id: "4", Name: "Ereignisfeld 1"}, func(playerId string) {
		cardEvent := board.DrawCard()
		cardEvent.PlayerId = playerId
	}},
	{propertyDomain.BaseFieldInformation{Id: "6", Name: "Ereignisfeld 2"}, func(playerId string) {
		fmt.Println("Ereignisfeld")
		//TODO implement ereignis
	}},
	{propertyDomain.BaseFieldInformation{Id: "11", Name: "Ereignisfeld 3"}, func(playerId string) {
		fmt.Printf("Remove 100 from Bank account of player %s\n", playerId)
		//TODO Add transaction
	}},
	{propertyDomain.BaseFieldInformation{Id: "13", Name: "Gehe ins Gefaengnis"}, func(playerId string) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		board.MovePlayer(playerId, 4)
	}},
	{propertyDomain.BaseFieldInformation{Id: "15", Name: "Ereignisfeld 4"}, func(playerId string) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
}

var defaultBasicFields = []propertyDomain.BasicField{
	{propertyDomain.BaseFieldInformation{Id: "1", Name: "Start"}},
	{propertyDomain.BaseFieldInformation{Id: "5", Name: "Gefaengnis"}},
	{propertyDomain.BaseFieldInformation{Id: "9", Name: "Frei parken"}},
}

var defaultProperties = []propertyDomain.PropertyField{
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "2", Name: "Property purple 1"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "3", Name: "Property purple 2"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "7", Name: "Property orange 1"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "8", Name: "Property orange 2"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "10", Name: "Property green 1"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "12", Name: "Property green 2"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "14", Name: "Property blue 1"}, tempFinancialDetails),
	*propertyDomain.NewPropertyField(propertyDomain.BaseFieldInformation{Id: "16", Name: "Property blue 2"}, tempFinancialDetails),
}

func initFields() {
	properties = nil
	basicFields = nil
	eventFields = nil
	properties = defaultProperties
	basicFields = defaultBasicFields
	eventFields = defaultEventFields
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

	for i := 0; i < len(defaultProperties); i++ {
		if defaultProperties[i].GetId() == fieldId {
			return &defaultProperties[i]
		}
	}
	panic("Field not found")
}

func GetFieldById(fieldId string) propertyDomain.Field {

	for i := 0; i < len(basicFields); i++ {
		if basicFields[i].GetId() == fieldId {
			return basicFields[i]
		}
	}
	for i := 0; i < len(eventFields); i++ {
		if eventFields[i].GetId() == fieldId {
			return eventFields[i]
		}
	}
	for i := 0; i < len(properties); i++ {
		if properties[i].GetId() == fieldId {
			return properties[i]
		}
	}
	panic("Field not found")
}

func GetFieldsCount() int {
	return len(properties) + len(basicFields) + len(eventFields)
}
