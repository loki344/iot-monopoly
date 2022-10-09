package eventing

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
)

func StartInternalEventBus() {
	fmt.Println("Starting eventbus")
	initChannels()
}

func initChannels() {

	//TODO external is simply to mock the rabbitMQ
	channels := []string{"external", "internal", "lapFinished"}
	tr := bus.GetBus()
	for _, channel := range channels {
		tr.GetChannelManager().CreateChannel(channel)
	}
}
