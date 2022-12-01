package eventing

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/eventing/config"
)

type ChannelName string

const (
	GAME_STARTED                 ChannelName = "gameStarted"
	PROPERTY_BUY_QUESTION        ChannelName = "propertyBuyQuestion"        // player lands on field and can buy it
	PROPERTY_TRANSACTION_STARTED ChannelName = "propertyTransactionStarted" // player is willing to buy a property, transaction process initiated
	TRANSACTION_REQUEST          ChannelName = "transactionAdded"           // player has to pay the amount with his card
	PAYMENT_REQUESTED            ChannelName = "paymentRequested"           // player lands on owned field and has to pay
	TRANSACTION_RESOLVED         ChannelName = "transactionResolved"        // transaction is finished and all money is moved
	LAP_FINISHED                 ChannelName = "lapFinished"                // player has finished a lap
	PAYOUT_REQUESTED             ChannelName = "payout"                     // payment request to payout money from bank to player
	CARD_DREW                    ChannelName = "cardDrew"                   // event card is drew
	MOVE_PLAYER                  ChannelName = "movePlayer"                 // player is moved by an event or go to prison field
)

func RegisterEventHandler(handler bus.Handler) string {
	fmt.Println("Registering handler for channel: " + handler.Matcher)
	createChannelIfNotExists(handler.Matcher)
	id := uuid.NewString()
	fmt.Printf("Handler received id %s\n", id)

	config.Bus.RegisterHandler(id, handler)
	return id
}

func DeregisterEventHandler(handlerId string) {
	fmt.Printf("Deregistering handler with id: %s \n", handlerId)
	config.Bus.DeregisterHandler(handlerId)
}

func FireEvent(channelName ChannelName, event any) {

	fmt.Println("Firing event:")
	fmt.Println(event)

	name := string(channelName)
	createChannelIfNotExists(name)

	txID := config.Monoton.Next()
	ctx := context.Background()
	ctx = context.WithValue(ctx, bus.CtxKeyTxID, txID)

	b := config.Bus

	err := b.Emit(
		ctx,
		name,
		event,
	)
	if err != nil {
		fmt.Println("ERROR >>>>", err)
	}
}

func createChannelIfNotExists(name string) {
	for _, topic := range config.Bus.Topics() {
		if topic == name {
			return
		}
	}

	fmt.Printf("Channel %s not found, creating it\n", name)
	config.Bus.RegisterTopics(name)
}
