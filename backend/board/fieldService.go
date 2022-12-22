package board

import (
	"fmt"
	"github.com/google/uuid"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
)

var properties []boardDomain.PropertyField
var pendingTransfer *boardDomain.PendingPropertyTransaction
var basicFields []boardDomain.BasicField
var eventFields []boardDomain.EventField

var tempFinancialDetails = &boardDomain.FinancialDetails{100, 100, 100, boardDomain.Revenue{100, 200, 300, 400, 500, 800}}

var defaultEventFields = []boardDomain.EventField{
	{boardDomain.BaseFieldInformation{Id: "4", Name: "Ereignisfeld 1"}, func(player *boardDomain.Player) {
		cardEvent := DrawCard()
		cardEvent.Player = player
	}},
	{boardDomain.BaseFieldInformation{Id: "6", Name: "Ereignisfeld 2"}, func(player *boardDomain.Player) {
		fmt.Println("Ereignisfeld")
		//TODO implement ereignis
	}},
	{boardDomain.BaseFieldInformation{Id: "11", Name: "Ereignisfeld 3"}, func(player *boardDomain.Player) {
		fmt.Printf("Remove 100 from Bank account of player %s\n", player.Id)
		//TODO Add transaction
	}},
	{boardDomain.BaseFieldInformation{Id: "13", Name: "Gehe ins Gefaengnis"}, func(player *boardDomain.Player) {
		fmt.Println("Player has to go to prison")
		// TODO this field index for prison should not be magic
		MovePlayer(player.Id, 4)
	}},
	{boardDomain.BaseFieldInformation{Id: "15", Name: "Ereignisfeld 4"}, func(player *boardDomain.Player) {
		//TODO implement ereignis
		fmt.Println("Ereignisfeld")
	}},
}

var defaultBasicFields = []boardDomain.BasicField{
	{boardDomain.BaseFieldInformation{Id: "1", Name: "Start"}},
	{boardDomain.BaseFieldInformation{Id: "5", Name: "Gefaengnis"}},
	{boardDomain.BaseFieldInformation{Id: "9", Name: "Frei parken"}},
}

var defaultProperties = []boardDomain.PropertyField{
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "2", Name: "Property purple 1"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "3", Name: "Property purple 2"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "7", Name: "Property orange 1"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "8", Name: "Property orange 2"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "10", Name: "Property green 1"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "12", Name: "Property green 2"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "14", Name: "Property blue 1"}, tempFinancialDetails),
	*boardDomain.NewPropertyField(boardDomain.BaseFieldInformation{Id: "16", Name: "Property blue 2"}, tempFinancialDetails),
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

	transactionId := uuid.NewString()
	property := *GetPropertyById(propertyId)
	eventing.FireEvent(eventing.PROPERTY_TRANSACTION_STARTED, boardDomain.NewTransactionRequestWithId(transactionId, "Bank", buyerId, property.GetPropertyFee()))
	pendingTransfer = &boardDomain.PendingPropertyTransaction{TransactionId: transactionId, PropertyId: propertyId, BuyerId: buyerId}

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

func GetPropertyById(fieldId string) *boardDomain.PropertyField {

	for i := 0; i < len(defaultProperties); i++ {
		if defaultProperties[i].GetId() == fieldId {
			return &defaultProperties[i]
		}
	}
	panic("Field not found")
}

func GetFieldById(fieldId string) boardDomain.Field {

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
