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
	PROPERTY_BUY_QUESTION ChannelName = "propertyBuyQuestion"
	TRANSACTION_ADDED     ChannelName = "transactionAdded"
	TRANSACTION_REQUESTED ChannelName = "transactionRequested"
	LAP_FINISHED          ChannelName = "lapFinished"
)

func RegisterEventHandler(handler bus.Handler) string {
	fmt.Println("Registering handler for channel: " + handler.Matcher)
	createChannelIfNotExists(handler.Matcher)
	id := uuid.NewString()
	fmt.Printf("Handler received id %s\n", id)

	config.Bus.RegisterHandler(id, handler)
	return id
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
