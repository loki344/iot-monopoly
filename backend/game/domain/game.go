package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/eventing"
	"math/rand"
	"time"
)

type Game struct {
	players            []Player
	currentPlayerIndex int
	ended              bool
	properties         []PropertyField
	eventFields        []EventField
	pendingTransfer    *PendingPropertyTransfer
	cards              []Card
	pendingCard        *Card
}

func (game *Game) Ended() bool {
	return game.ended
}

func (game *Game) Players() []Player {
	return game.players
}

func (game *Game) CurrentPlayerIndex() int {
	return game.currentPlayerIndex
}

var defaultFinancialDetails = FinancialDetails{100, 100, 100, Revenue{100, 200, 300, 400, 500, 800}}

//var defaultBasicFields = []domain.BasicField{
//	{domain.BaseFieldInformation{index: "1", name: "Start"}},
//	{domain.BaseFieldInformation{index: "5", name: "Gefaengnis"}},
//	{domain.BaseFieldInformation{index: "9", name: "Frei parken"}},
//}

var defaultProperties = []PropertyField{
	*NewPropertyField("Property purple 1", 2, defaultFinancialDetails),
	*NewPropertyField("Property purple 2", 3, defaultFinancialDetails),
	*NewPropertyField("Property orange 1", 7, defaultFinancialDetails),
	*NewPropertyField("Property orange 2", 8, defaultFinancialDetails),
	*NewPropertyField("Property green 1", 10, defaultFinancialDetails),
	*NewPropertyField("Property green 2", 12, defaultFinancialDetails),
	*NewPropertyField("Property blue 1", 14, defaultFinancialDetails),
	*NewPropertyField("Property blue 2", 16, defaultFinancialDetails),
}
var defaultEventFields = []EventField{
	*NewEventField("Ereignisfeld 1", 4, DRAW_CARD),
	*NewEventField("Ereignisfeld 2", 6, DRAW_CARD),
	*NewEventField("Einkommenssteuer", 11, PAY_TAX),
	*NewEventField("Gehe ins Gefaengnis", 13, GOTO_PRISON),
	*NewEventField("Ereignisfeld 4", 15, DRAW_CARD),
}

var defaultCardStack = []Card{
	*NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(player *Player) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED, NewGameEventWithPayoutAcceptedEvent(player.Account().Id(), 100))
	}),
	*NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(player *Player) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, NewGameEventWithFeeAcceptedEvent("Bank", player.Account().Id(), 200))
	}),
}

func NewGame(playerCount int) *Game {

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]Player, playerCount)

	for i := 0; i < playerCount; i++ {
		newPlayers[i] = *NewPlayer(i + 1)
	}

	newPlayers = append(newPlayers, *CreateBank())

	eventing.FireEvent(eventing.GAME_STARTED, NewGameStartedEvent(playerCount))
	return &Game{players: newPlayers, currentPlayerIndex: 0, properties: defaultProperties, eventFields: defaultEventFields, cards: defaultCardStack}
}

func (game *Game) TransferOwnership(transactionId string) {
	pendingTransfer := game.pendingTransfer

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.Id() == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyIndex(), pendingTransfer.BuyerId())
		property := game.GetPropertyByIndex(pendingTransfer.PropertyIndex())
		property.ownerId = pendingTransfer.BuyerId()
		game.pendingTransfer = nil
	}
}

func (game *Game) End(winnerId string) {
	eventing.FireEvent(eventing.GAME_ENDED, NewGameEndedEvent(winnerId))
}

func (game Game) PlayerCount() int {
	return len(game.players)
}

func (game *Game) MovePlayer(playerId string, position int) error {

	player := game.GetPlayerById(playerId)

	//TODO move this code somewhere where it makes more sense
	totalFieldCount := 16
	if position > totalFieldCount-1 || position < 0 {
		return errors.New(fmt.Sprintf("Fieldindex %d out of bound for Fieldlength %d", position, totalFieldCount))
	}

	game.updateCurrentPlayerIndex()
	game.checkIfLapFinished(playerId, position, player)
	player.SetPosition(position)

	if game.GetPropertyByIndex(position) != nil {

		property := game.GetPropertyByIndex(position)
		if property.ownerId == "" {
			fmt.Println("property has no owner, is buyable")
			eventing.FireEvent(eventing.PLAYER_ON_UNOWNED_FIELD, NewPlayerOnUnownedFieldEvent(playerId, property))
		} else if property.ownerId != playerId {
			fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", property.ownerId, playerId, property.GetPropertyFee())
			// Create transaction
		}
	}

	if game.GetEventFieldByIndex(position) != nil {
		eventField := game.GetEventFieldByIndex(position)
		switch eventField.Type() {

		case DRAW_CARD:
			game.drawCard(playerId)
			break
		case GOTO_PRISON:
			//TODO implement
			break
		case PAY_TAX:
			eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, NewGameEventWithFeeAcceptedEvent("Bank", game.GetPlayerById(playerId).Account().Id(), 200))
			break
		}
	}

	return nil
}

func (game *Game) drawCard(playerId string) {
	rand.Seed(time.Now().UnixNano())
	card := game.cards[rand.Intn(len(game.cards))]
	card.SetPlayer(*game.GetPlayerById(playerId))

	game.pendingCard = &card
	eventing.FireEvent(eventing.CARD_DREW, NewCardDrewEvent(card.title, card.text))
}

func (game *Game) checkIfLapFinished(playerId string, position int, player *Player) {
	if player.Position() > position && position < 5 {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, NewLapFinishedEvent(playerId))
		player.Account().Deposit(100)
	}
}

func (game *Game) updateCurrentPlayerIndex() {
	if game.currentPlayerIndex == len(game.players)-1 {
		game.currentPlayerIndex = 0
	} else {
		game.currentPlayerIndex = game.currentPlayerIndex + 1
	}
}

func (game *Game) GetPlayerById(playerId string) *Player {

	for i := range game.players {
		if game.players[i].Id() == playerId {
			return &game.players[i]
		}
	}

	panic(fmt.Sprintf("Player with index %s not found", playerId))
}

func (game *Game) GetPropertyByIndex(index int) *PropertyField {

	for i := range game.properties {
		if game.properties[i].index == index {
			return &game.properties[i]
		}
	}
	return nil
}
func (game *Game) GetEventFieldByIndex(index int) *EventField {

	for i := range game.eventFields {
		if game.eventFields[i].index == index {
			return &game.eventFields[i]
		}
	}
	return nil
}
func (game *Game) FindAccountById(accountId string) *Account {

	for i := range game.players {
		if game.players[i].Account().Id() == accountId {
			return game.players[i].Account()
		}
	}
	return nil
}

func (game *Game) BuyProperty(propertyIndex int, buyerId string) string {

	game.pendingTransfer = NewPendingPropertyTransaction(uuid.NewString(), propertyIndex, buyerId)
	return game.pendingTransfer.id
}

func (game *Game) ConfirmCurrentCard() {
	game.pendingCard.TriggerAction()
}
