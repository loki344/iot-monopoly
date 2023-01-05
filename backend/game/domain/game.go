package domain

import (
	"fmt"
	"github.com/google/uuid"
	"iot-monopoly/eventing"
	"iot-monopoly/game/domain/events"
	"math/rand"
	"time"
)

type Game struct {
	players            []Player
	currentPlayerIndex int
	ended              bool
	board              *Board
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

func (game Game) PlayerCount() int {
	return len(game.players)
}

func NewGame(playerCount int) *Game {

	fmt.Printf("starting game with %d players\n", playerCount)
	newPlayers := make([]Player, playerCount)

	for i := 0; i < playerCount; i++ {
		newPlayers[i] = *NewPlayer(i + 1)
	}

	newPlayers = append(newPlayers, *CreateBank())

	eventing.FireEvent(eventing.GAME_STARTED, events.NewGameStartedEvent(playerCount))
	return &Game{players: newPlayers, currentPlayerIndex: 0, board: NewBoard(defaultProperties, defaultEventFields), cards: defaultCardStack}
}

func (game *Game) TransferOwnership(transactionId string) {
	pendingTransfer := game.pendingTransfer

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.Id() == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyIndex(), pendingTransfer.BuyerId())
		property := game.board.GetPropertyByIndex(pendingTransfer.PropertyIndex())
		property.ownerId = pendingTransfer.BuyerId()
		game.pendingTransfer = nil
	}
}

func (game *Game) End(winnerId string) {
	eventing.FireEvent(eventing.GAME_ENDED, events.NewGameEndedEvent(winnerId))
}

func (game *Game) MovePlayer(playerId string, position int) {

	player := game.GetPlayerById(playerId)
	if player.Position() == position {
		fmt.Println(fmt.Errorf("player already at position %d", position))
		return
	}

	game.updateCurrentPlayerIndex()
	game.checkIfLapFinished(playerId, position, player)
	player.position = position

	game.triggerFieldAction(playerId, position)
}

func (game *Game) triggerFieldAction(playerId string, position int) {

	if game.board.GetPropertyByIndex(position) != nil {

		property := game.board.GetPropertyByIndex(position)
		if property.ownerId == "" {
			fmt.Println("property has no owner, is buyable")
			eventing.FireEvent(eventing.PLAYER_ON_UNOWNED_FIELD, events.NewPlayerOnUnownedFieldEvent(playerId, property))
		} else if property.ownerId != playerId {
			fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", property.ownerId, playerId, property.GetPropertyFee())
			// Create transaction
		}
	}

	if game.board.GetEventFieldByIndex(position) != nil {
		eventField := game.board.GetEventFieldByIndex(position)
		switch eventField.Type() {

		case DRAW_CARD:
			game.drawCard(playerId)
			break
		case GOTO_PRISON:
			//TODO implement
			break
		case PAY_TAX:
			eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, events.NewGameEventWithFeeAcceptedEvent("Bank", game.GetPlayerById(playerId).Account().Id(), 200))
			break
		}
	}
}

func (game *Game) drawCard(playerId string) {
	rand.Seed(time.Now().UnixNano())
	card := game.cards[rand.Intn(len(game.cards))]
	card.SetPlayer(*game.GetPlayerById(playerId))

	game.pendingCard = &card
	eventing.FireEvent(eventing.CARD_DREW, events.NewCardDrewEvent(card.title, card.text))
}

func (game *Game) checkIfLapFinished(playerId string, position int, player *Player) {
	if player.Position() > position && position < 5 {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, events.NewLapFinishedEvent(playerId))
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

func (game *Game) BuyProperty(propertyIndex int, buyerId string) string {

	game.pendingTransfer = NewPendingPropertyTransaction(uuid.NewString(), propertyIndex, buyerId)
	return game.pendingTransfer.id
}

func (game *Game) ConfirmCurrentCard() {
	game.pendingCard.TriggerAction()
}
