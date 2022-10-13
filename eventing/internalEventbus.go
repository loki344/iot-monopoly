package eventing

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
)

type ChannelName string

const (
	EXTERNAL              ChannelName = "external"
	PROPERTY_BUY_QUESTION ChannelName = "propertyBuyQuestion"
	TRANSACTION_REQUESTED ChannelName = "transactionRequested"
	LAP_FINISHED          ChannelName = "lapFinished"
)

func ListenRequestStream(channelName ChannelName) bus.MessageHandler {

	createChannelIfNotExists(channelName)
	tr := bus.GetBus()

	eventHandler, err := tr.ListenRequestStream(string(channelName))
	if err != nil {
		fmt.Println(err)
	}

	return eventHandler
}

func FireEvent(channelName ChannelName, event any) {

	name := string(channelName)
	createChannelIfNotExists(channelName)
	tr := bus.GetBus()
	handler, err := tr.RequestOnce(name, event)
	if err != nil {
		//TODO something went wrong
		fmt.Println(err)
	}
	err = handler.Fire()
	if err != nil {
		//TODO something went wrong
		fmt.Println(err)
	}
}

func createChannelIfNotExists(name ChannelName) {
	tr := bus.GetBus()

	if !tr.GetChannelManager().CheckChannelExists(string(name)) {
		fmt.Printf("Channel %s not found, creating it\n", name)
		tr.GetChannelManager().CreateChannel(string(name))
	}
}
