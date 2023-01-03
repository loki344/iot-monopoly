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
	GAME_STARTED                    ChannelName = "gameStarted"
	GAME_ENDED                      ChannelName = "gameEnded"
	PLAYER_MOVED                    ChannelName = "playerMoved"
	PLAYER_ON_UNOWNED_FIELD         ChannelName = "playerOnUnownedField"
	PLAYER_ON_OWNED_FIELD           ChannelName = "playerOnOwnedField"
	PROPERTY_TRANSFER_CREATED       ChannelName = "propertyTransferCreated"
	TRANSACTION_CREATED             ChannelName = "transactionCreated"
	TRANSACTION_RESOLVED            ChannelName = "transactionResolved"
	LAP_FINISHED                    ChannelName = "lapFinished"
	GAME_EVENT_WITH_PAYOUT_ACCEPTED ChannelName = "gameEventWithPayoutAccepted"
	GAME_EVENT_WITH_FEE_ACCEPTED    ChannelName = "gameEventWithFeeAccepted"
	CARD_DREW                       ChannelName = "cardDrew"
	PLAYER_BANKRUPT                 ChannelName = "playerMoved"
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
