package domain

import (
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
	board              *Board
	pendingTransfer    *PendingPropertyTransfer
	cards              []Card
	pendingCard        *Card
	bank               *Bank
}

func (game *Game) Bank() *Bank {
	return game.bank
}

func (game *Game) Board() *Board {
	return game.board
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
		newPlayers[i] = *newPlayer(i + 1)
	}

	eventing.FireEvent(eventing.GAME_STARTED, newGameStartedEvent(playerCount))
	return &Game{players: newPlayers, bank: newBank(), currentPlayerIndex: 0, board: newBoard(defaultProperties, defaultEventFields, standardFields, newPrison()), cards: defaultCardStack}
}

func (game *Game) ResolvePendingPropertyTransfer(transactionId string) {
	pendingTransfer := game.pendingTransfer

	if pendingTransfer == nil {
		return
	}

	if pendingTransfer.ReferenceTransactionId() == transactionId {
		fmt.Printf("Transferring ownership for property %s to %s\n", pendingTransfer.PropertyIndex(), pendingTransfer.BuyerId())
		property := game.board.GetPropertyByIndex(pendingTransfer.PropertyIndex())
		property.ownerId = pendingTransfer.BuyerId()
		game.pendingTransfer = nil
	}
}

func (game *Game) End(winnerId string) {
	eventing.FireEvent(eventing.GAME_ENDED, newGameEndedEvent(winnerId))
}

func (game *Game) MovePlayer(playerId string, newPosition int) {

	player := game.GetPlayerById(playerId)
	if player.Position() == newPosition {
		fmt.Println(fmt.Errorf("player already at position %d", newPosition))
		return
	}

	game.depositMoneyIfLapFinished(newPosition, player)
	player.position = newPosition
	game.triggerFieldAction(playerId, newPosition)
	game.updateCurrentPlayerIndex()
}

func (game *Game) triggerFieldAction(playerId string, position int) {

	if game.board.GetPropertyByIndex(position) != nil {

		property := game.board.GetPropertyByIndex(position)
		if property.ownerId == "" {
			fmt.Println("property has no owner, is buyable")
			eventing.FireEvent(eventing.PLAYER_ON_UNOWNED_FIELD, newPlayerOnUnownedFieldEvent(playerId, property))
		} else if property.ownerId != playerId {
			fmt.Printf("Property belongs to player %s, player %s has to pay %d\n", property.ownerId, playerId, property.GetPropertyFee())
			senderAccountId := game.GetPlayerById(playerId).Account().Id()
			recipientAccountId := game.GetPlayerById(property.ownerId).Account().Id()
			eventing.FireEvent(eventing.PLAYER_ON_OWNED_FIELD, newPlayerOnOwnedFieldEvent(senderAccountId, recipientAccountId, property.GetPropertyFee()))
		}
	}

	if game.board.GetEventFieldByIndex(position) != nil {
		eventField := game.board.GetEventFieldByIndex(position)
		switch eventField.Type() {
		case DRAW_CARD:
			game.drawCard(playerId)
			break
		case GOTO_PRISON:
			game.goToPrison(playerId)
			break
		case PAY_TAX:
			eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, newGameEventWithFeeAcceptedEvent("Bank", game.GetPlayerById(playerId).Account().Id(), 200))
			break
		}
	}
}

func (game *Game) drawCard(playerId string) {
	rand.Seed(time.Now().UnixNano())
	card := game.cards[rand.Intn(len(game.cards))]
	card.SetPlayer(game.GetPlayerById(playerId))

	game.pendingCard = &card
	eventing.FireEvent(eventing.CARD_DREW, newCardDrewEvent(card.title, card.text))
}

func (game *Game) depositMoneyIfLapFinished(position int, player *Player) {
	if player.Position() > position && position < 5 {
		fmt.Println("Player completed a lap, creating lap finished")
		eventing.FireEvent(eventing.LAP_FINISHED, newLapFinishedEvent(player.id))
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

	game.pendingTransfer = newPendingPropertyTransaction(uuid.NewString(), propertyIndex, buyerId)
	return game.pendingTransfer.referenceTransactionId
}

func (game *Game) ConfirmCurrentCard() {
	game.pendingCard.TriggerAction()
}

func (game *Game) goToPrison(playerId string) {
	game.board.goToPrison(playerId)
	game.GetPlayerById(playerId).position = game.board.prison.index
}

func (game *Game) Properties() []PropertyField {
	return game.board.properties
}
