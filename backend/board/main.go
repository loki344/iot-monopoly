package board

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
	"strconv"
)

var players []boardDomain.Player

//TODO use static uuid
var defaultPlayers = []boardDomain.Player{
	{"Player_1", 0, "Account_Player_1"},
	{"Player_2", 0, "Account_Player_2"},
	{"Player_3", 0, "Account_Player_3"},
	{"Player_4", 0, "Account_Player_4"},
	//33-A8-8A-10 Card 1
	//	1304-B6-1A Card 2
	// 43-F1-E70E Card 3
	// A3-D9-350F Card 4
}

var tempFinancialDetails = &boardDomain.FinancialDetails{100, 100, 100, boardDomain.Revenue{100, 200, 300, 400, 500, 800}}

var defaultEventFields = []boardDomain.EventField{
	{boardDomain.BaseFieldInformation{Id: "4", Name: "Ereignisfeld 1"}, func(player *boardDomain.Player) {
		fmt.Println("Ereignisfeld")
		//TODO implement ereignis
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
	{boardDomain.BaseFieldInformation{Id: "15", Name: "Ereignisfeld 3"}, func(player *boardDomain.Player) {
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

var properties []boardDomain.PropertyField
var basicFields []boardDomain.BasicField
var eventFields []boardDomain.EventField

var pendingTransfer boardDomain.PendingPropertyTransaction

func StartGame(playerCount int) ([]boardDomain.Player, error) {
	eventing.FireEvent(eventing.GAME_STARTED, boardDomain.NewGameStartedEvent(playerCount))
	players = nil
	properties = nil
	basicFields = nil
	eventFields = nil
	properties = defaultProperties
	basicFields = defaultBasicFields
	eventFields = defaultEventFields

	if playerCount < 1 || playerCount > 4 {
		errorMsg := fmt.Sprintf("invalid playerCount %d (must be between 1 and 4)", playerCount)
		fmt.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]boardDomain.Player, playerCount)

	copy(newPlayers, defaultPlayers)

	players = newPlayers

	return players, nil
}

func MovePlayer(playerId string, fieldId int) error {

	totalFieldCount := GetFieldsCount()
	if fieldId > totalFieldCount-1 || fieldId < 0 {
		return errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", fieldId, totalFieldCount))
	}

	player := GetPlayer(playerId)
	if player.Position == fieldId {
		fmt.Println(fmt.Errorf("player already at position %d", fieldId))
		return nil
	}

	//TODO get rid of magic numbers 10!!
	if (player.Position >= 10 && player.Position < totalFieldCount) && (fieldId >= 0 && fieldId <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, boardDomain.NewLapFinishedEvent(player.Id))
	}

	fmt.Printf("MovePlayer player %s to fieldId %d\n", player.Id, fieldId)
	player.Position = fieldId
	GetFieldById(strconv.FormatInt(int64(fieldId), 10)).OnPlayerEnter(player)
	return nil
}

func BuyProperty(propertyId string, buyerId string) string {

	transactionId := uuid.NewString()
	property := *GetPropertyById(propertyId)
	eventing.FireEvent(eventing.PROPERTY_TRANSACTION_STARTED, boardDomain.NewTransactionRequestWithId(transactionId, "Bank", buyerId, property.GetPropertyFee()))
	pendingTransfer = boardDomain.PendingPropertyTransaction{TransactionId: transactionId, PropertyId: propertyId, BuyerId: buyerId}

	return transactionId
}

func transferOwnerShip(transactionId string) {

	if pendingTransfer.TransactionId == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyId, pendingTransfer.BuyerId)
		property := GetPropertyById(pendingTransfer.PropertyId)
		property.OwnerId = pendingTransfer.BuyerId
		return
	}
}

func GetPlayer(playerId string) *boardDomain.Player {

	for i := range players {
		if players[i].Id == playerId {
			return &players[i]
		}
	}

	panic(fmt.Sprintf("Player with id %s not found", playerId))
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

func GetPlayers() []boardDomain.Player {
	return players
}
