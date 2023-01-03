package repository

import "iot-monopoly/property/domain"

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

func InitFields() {
	properties = nil
	properties = defaultProperties
}

func CreatePendingPropertyTransaction(transaction *domain.PendingPropertyTransaction) {
	pendingTransfer = transaction
}

func DeletePendingPropertyTransaction(transaction *domain.PendingPropertyTransaction) {
	if pendingTransfer.Id() == transaction.Id() {
		pendingTransfer = nil
	}
}

func FindPropertyById(fieldId string) *domain.PropertyField {

	for i := 0; i < len(properties); i++ {
		if properties[i].Id == fieldId {
			return properties[i]
		}
	}
	return nil
}

func GetPendingTransfer() *domain.PendingPropertyTransaction {
	return pendingTransfer
}
